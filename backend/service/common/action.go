package common

import (
	"PeachDRAC/backend/constants"
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/service/control"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"fmt"
	"sync"
)

func (s *CommonService) Action(request model.ActionRequest) {

	if len(request.IPs) == 0 {
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(len(request.IPs))

	// 获取可用的密码组
	passwdGroups := s.ConfigPasswd.GetAll()

	for _, ip := range request.IPs {
		go func(ip string) {
			defer wg.Done()

			if !passwdGroups.Status {
				runtime.EventsEmit(s.Ctx, "actions", model.ActionRespond{
					IP:     ip,
					Status: false,
					Action: request.Action,
					Result: fmt.Sprintf("获取密码组失败: %s", passwdGroups.Msg),
				})
				return
			}
			if len(passwdGroups.Data.([]model.Passwd)) == 0 {
				runtime.EventsEmit(s.Ctx, "actions", model.ActionRespond{
					IP:     ip,
					Status: false,
					Action: request.Action,
					Result: "密码组为空",
				})
				return
			}

			var (
				isSuccess   = false                 // 是否成功
				ipmiControl *control.ControlService // 控制服务
			)

			// 遍历密码组
			for _, v := range passwdGroups.Data.([]model.Passwd) {
				// 实例化
				ipmiControl = control.NewService(ip, v.Username, v.Password, v.Port)
				// 连接
				if err := ipmiControl.ConnectServer(); err != nil {
					continue
				}
				isSuccess = true // 登录成功
				break
			}

			// 连接失败
			if !isSuccess {
				runtime.EventsEmit(s.Ctx, "actions", model.ActionRespond{
					IP:     ip,
					Status: false,
					Action: request.Action,
					Result: fmt.Sprintf("已尝试%d个密码组,均登录失败", len(passwdGroups.Data.([]model.Passwd))),
				})
				return
			}

			// 获取设备型号
			system, err := ipmiControl.GetSystem()
			if err != nil {
				runtime.EventsEmit(s.Ctx, "actions", model.ActionRespond{
					IP:     ip,
					Status: false,
					Action: request.Action,
					Result: fmt.Sprintf("获取设备型号失败: %s", err.Error()),
				})
				return
			}

			// 执行对应的操作
			var errSwitch error

			switch request.Action {
			case constants.ActionPowerOn:
				errSwitch = ipmiControl.PowerOn() // 开机
			case constants.ActionPowerOff:
				errSwitch = ipmiControl.PowerOff() // 关机
			case constants.ActionPowerReset:
				errSwitch = ipmiControl.PowerRestart() // 重启
			case constants.ActionPowerForceReset:
				errSwitch = ipmiControl.PowerHardRestart() // 强制重启
			case constants.ActionFanAdjust:
				errSwitch = ipmiControl.FanAdjust(uint8(request.Fan.Speed)) // 风扇调整
			case constants.ActionFanAdaptive:
				errSwitch = ipmiControl.FanAdaptive() // 风扇自适应
			case constants.ActionMountNFS:
			}

			if errSwitch != nil {
				runtime.EventsEmit(s.Ctx, "actions", model.ActionRespond{
					IP:          ip,
					Status:      false,
					Action:      request.Action,
					ProductName: system.ProductName,
					Result:      fmt.Sprintf("执行指令失败: %s", errSwitch.Error()),
				})
			}

			runtime.EventsEmit(s.Ctx, "actions", model.ActionRespond{
				IP:          ip,
				Status:      true,
				Action:      request.Action,
				Result:      "执行成功",
				ProductName: system.ProductName,
			})

			defer ipmiControl.Close()

		}(ip)
	}
	wg.Wait()
}
