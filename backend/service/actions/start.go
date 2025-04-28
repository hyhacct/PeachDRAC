package actions

import (
	"PeachDRAC/backend/constants"
	"PeachDRAC/backend/farmework"
	interfaces "PeachDRAC/backend/interfaces"
	"PeachDRAC/backend/model"
	"fmt"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (s *ServiceActions) Start(ips []string, action string, fan int, nfs string) model.WailsCommunicate {

	// 是否为空 IP 组
	if len(ips) == 0 {
		return model.WailsError("没有IP地址")
	}

	// 得到所有密码组
	pass_list, err := model.TablePass{}.GetAllEnabled()
	if err != nil {
		return model.WailsError("获取密码组失败")
	}

	if len(pass_list) == 0 {
		return model.WailsError("没有密码组")
	}

	wg := sync.WaitGroup{}
	for _, ip := range ips {
		wg.Add(1)
		go func(ip string) {
			var (
				client    = &interfaces.InterfacesDefault{} // IPMI客户端
				isSuccess = false                           // 是否成功
			)

			for _, item := range pass_list {
				if err := client.Connect(ip, item.Username, item.Password, item.Port); err != nil {
					continue
				}
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

			if !isSuccess {
				runtime.EventsEmit(s.ctx, constants.EventActions, model.WailsActionsError(client.Actions, fmt.Sprintf("已尝试%d个密码组,均失败", len(pass_list))))
				return
			}

			// 读取硬件信息(最多尝试3次)
			for i := 0; i < 3; i++ {
				deviceModel, sn, manufacturer, err := client.GetModelAndSN()
				if err != nil {
					continue
				}
				client.Actions.DeviceModel = deviceModel
				client.Actions.Sn = sn
				client.Actions.Manufacturer = manufacturer
				break
			}
			if err != nil {
				runtime.EventsEmit(s.ctx, constants.EventActions, model.WailsActionsError(client.Actions, fmt.Sprintf("读取硬件信息失败: %s", err)))
				return
			}

			// 记录一下日志
			farmework.ModuleLogs.Info("进行操作", client.Actions, "服务器=", ip, "SN=", client.Actions.Sn, "型号=", client.Actions.DeviceModel, "厂商=", client.Actions.Manufacturer)

			// 根据不同的厂商执行对应的接口
			switch client.Actions.Manufacturer {
			case "DELL":
				s.BranchDell(client)
			case "Inspur":
				s.BranchInspur(client)
			default:
				runtime.EventsEmit(s.ctx, constants.EventActions, model.WailsActionsError(client.Actions, fmt.Sprintf("暂不支持厂商:%s", client.Actions.Manufacturer)))
				return
			}

			// 释放资源
			defer func() {
				wg.Done()
				client.Close()
			}()
		}(ip)
	}

	// 等待所有探测完成
	wg.Wait()

	return model.WailsSuccess("探测完成", "")

}
