package actions

import (
	"PeachDRAC/backend/constants"
	"PeachDRAC/backend/encapsulation"
	"PeachDRAC/backend/farmework"
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
			defer wg.Done()

			var (
				client    = &encapsulation.IPMI{} // IPMI客户端
				isSuccess = false                 // 是否成功
				actions   model.ModelActions      // 操作数据
			)

			for _, item := range pass_list {
				if err := client.Connect(ip, item.Username, item.Password, item.Port); err != nil {
					continue
				}
				isSuccess = true
				actions = model.ModelActions{
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
				runtime.EventsEmit(s.ctx, constants.EventActions, model.WailsActionsError(ip, action, fmt.Sprintf("已尝试%d个密码组,均失败", len(pass_list))))
				return
			}

			// 读取硬件信息(最多尝试3次)
			for i := 0; i < 3; i++ {
				deviceModel, sn, manufacturer, err := client.GetModelAndSN()
				if err != nil {
					continue
				}
				actions.DeviceModel = deviceModel
				actions.Sn = sn
				actions.Manufacturer = manufacturer
				break
			}
			if err != nil {
				runtime.EventsEmit(s.ctx, constants.EventActions, model.WailsActionsError(ip, action, fmt.Sprintf("读取硬件信息失败: %s", err)))
				return
			}

			// 记录一下日志
			farmework.ModuleLogs.Info("进行操作", actions, "服务器=", ip, "SN=", actions.Sn, "型号=", actions.DeviceModel, "厂商=", actions.Manufacturer)

			// 根据不同的厂商执行对应的接口
			switch actions.Manufacturer {
			case "DELL":
				s.ActionsDell(actions)
			}
		}(ip)
	}

	// 等待所有探测完成
	wg.Wait()

	return model.WailsSuccess("探测完成", "")

}

func (s *ServiceActions) ActionsDell(actions model.ModelActions) {
	var (
		api = &encapsulation.RedfishDell{
			IP:       actions.IP,
			Username: actions.Username,
			Password: actions.Password,
		}
		err error
	)
	switch actions.Action {
	case constants.ActionPowerOn:
		err = api.PowerOn() // 开机
	case constants.ActionRestart:
		err = api.ForceRestart() // 重启
	case constants.ActionPowerOff:
		err = api.ForceOff() // 关机
	case constants.ActionMountNFS:
		err = api.MountNFS(actions.Nfs) // 挂载NFS
	case constants.ActionUnmountNFS:
		err = api.UnmountNFS() // 卸载NFS
	case constants.ActionStartJavaConsole:
		err = api.LoginWebIpmiR730() // 登录Web Ipmi R730
		if err == nil {
			err = api.DownloadJnlp() // 下载JNLP文件
		}
	default:
		runtime.EventsEmit(s.ctx, constants.EventActions, model.WailsActionsError(actions.IP, actions.Action, "暂不支持此操作"))
		return
	}
	// 最终检查有没有错误
	if err != nil {
		runtime.EventsEmit(s.ctx, constants.EventActions, model.WailsActionsError(actions.IP, actions.Action, fmt.Sprintf("检查出错: %s", err.Error())))
		return
	}
	runtime.EventsEmit(s.ctx, constants.EventActions, model.WailsActionsSuccess(actions.IP, actions.Action, actions.DeviceModel, actions.Manufacturer, actions.Sn))
}
