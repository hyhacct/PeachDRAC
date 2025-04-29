package actions

import (
	"PeachDRAC/backend/constants"
	"PeachDRAC/backend/farmework"
	interfaces "PeachDRAC/backend/interfaces"
	"PeachDRAC/backend/model"
	"context"
	"fmt"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (s *ServiceActions) Start(ips []string, action string, fan int, nfs string) model.WailsCommunicate {
	if len(ips) == 0 {
		return model.WailsError("没有IP地址")
	}

	passList, err := model.TablePass{}.GetAllEnabled()
	if err != nil {
		return model.WailsError("获取密码组失败")
	}
	if len(passList) == 0 {
		return model.WailsError("没有密码组")
	}

	s.r = farmework.NewRunner()
	s.r.Run(len(ips))

	for _, ip := range ips {
		ip := ip // 避免闭包变量捕获问题
		s.r.Submit(farmework.TaskFunc(func(ctx context.Context) error {
			client := &interfaces.InterfacesDefault{}
			var isSuccess bool

			for _, item := range passList {
				if ctx.Err() != nil {
					return ctx.Err()
				}
				if err := client.Connect(ip, item.Username, item.Password, item.Port); err == nil {
					isSuccess = true
					client.Actions = model.ModelActions{
						IP:       ip,
						Username: item.Username,
						Password: item.Password,
						Action:   action,
						Fan:      fan,
						Nfs:      nfs,
					}
					break
				}
			}

			if !isSuccess {
				runtime.EventsEmit(s.ctx, constants.EventActions, model.WailsActionsError(client.Actions, fmt.Sprintf("已尝试%d个密码组,均失败", len(passList))))
				return fmt.Errorf("IP=%s 密码尝试失败", ip)
			}

			var err error
			for i := 0; i < 3; i++ {
				if ctx.Err() != nil {
					return ctx.Err()
				}
				var modelStr, sn, mfr string
				modelStr, sn, mfr, err = client.GetModelAndSN()
				if err == nil {
					client.Actions.DeviceModel = strings.TrimSpace(modelStr)
					client.Actions.Sn = strings.TrimSpace(sn)
					client.Actions.Manufacturer = strings.TrimSpace(mfr)
					break
				}
			}
			if err != nil {
				runtime.EventsEmit(s.ctx, constants.EventActions, model.WailsActionsError(client.Actions, fmt.Sprintf("读取硬件信息失败: %s", err)))
				return err
			}

			farmework.ModuleLogs.Info("进行操作", client.Actions, "服务器=", ip, "SN=", client.Actions.Sn, "型号=", client.Actions.DeviceModel, "厂商=", client.Actions.Manufacturer)

			if ctx.Err() != nil {
				return ctx.Err()
			}

			switch client.Actions.Manufacturer {
			case "DELL":
				s.BranchDell(client)
			case "Inspur":
				s.BranchInspur(client)
			default:
				runtime.EventsEmit(s.ctx, constants.EventActions, model.WailsActionsError(client.Actions, fmt.Sprintf("暂不支持厂商:%s", client.Actions.Manufacturer)))
				return fmt.Errorf("不支持的厂商: %s", client.Actions.Manufacturer)
			}

			defer client.Close()
			return nil
		}))
	}

	// 阻塞等待所有任务完成
	if err := s.r.Wait(); err != nil {
		return model.WailsError("部分任务失败")
	}

	return model.WailsSuccess("批量动作完成", "")
}
