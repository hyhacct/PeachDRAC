package common

import (
	"PeachDRAC/backend/constants"
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/service/control"
	"PeachDRAC/backend/service/redfish"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"fmt"
	"sync"
)

// 操作类变量集合
type operate struct {
	ipmiControl *control.ControlService
	ip          string
	username    string
	password    string
	port        int
	request     *model.ActionRequest // 请求参数
}

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

			// 实例化操作类
			var operate = &operate{}

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
				operate.username = v.Username
				operate.password = v.Password
				operate.port = v.Port
				operate.ipmiControl = ipmiControl
				operate.request = &request
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

			if system.Manufacturer == "Dell" {
				errSwitch = s.handleDellAction(operate)
			} else if system.Manufacturer == "Inspur" {
				errSwitch = s.handleInspurAction(operate)
			}

			if errSwitch != nil {
				runtime.EventsEmit(s.Ctx, "actions", model.ActionRespond{
					IP:          ip,
					Status:      false,
					Action:      request.Action,
					ProductName: system.ProductName,
					Result:      fmt.Sprintf("执行指令失败: %s", errSwitch.Error()),
				})
			} else {
				runtime.EventsEmit(s.Ctx, "actions", model.ActionRespond{
					IP:          ip,
					Status:      true,
					Action:      request.Action,
					Result:      "执行成功",
					ProductName: system.ProductName,
				})
			}

			defer ipmiControl.Close()

		}(ip)
	}
	wg.Wait()
}

func (s *CommonService) handleDellAction(operate *operate) error {
	switch operate.request.Action {
	case constants.ActionPowerOn:
		return operate.ipmiControl.DellPowerOn() // 开机
	case constants.ActionPowerOff:
		return operate.ipmiControl.DellPowerOff() // 关机
	case constants.ActionPowerReset:
		return operate.ipmiControl.DellPowerRestart() // 重启
	case constants.ActionPowerForceReset:
		return operate.ipmiControl.DellPowerHardRestart() // 硬重启
	case constants.ActionFanAdaptive:
		return operate.ipmiControl.DellFanAdaptive() // 风扇自适应
	case constants.ActionFanAdjust:
		return operate.ipmiControl.DellFanAdjust(uint8(operate.request.Fan.Speed)) // 风扇调节
	case constants.ActionUnmountNFS:
		return redfish.DellUmountNFS(operate.ip, operate.username, operate.password) // 卸载NFS
	case constants.ActionMountNFS:
		return redfish.DellMountNFS(operate.ip, operate.username, operate.password, fmt.Sprintf("%s:%s", operate.request.NFS.Mount.IP, operate.request.NFS.Mount.Path)) // 挂载NFS
	}
	return fmt.Errorf("戴尔暂不支持操作: %s", operate.request.Action)
}

func (s *CommonService) handleInspurAction(operate *operate) error {
	switch operate.request.Action {
	case constants.ActionMountNFS:
		return redfish.InspurMountNFS(operate.ip, operate.username, operate.password, operate.request.NFS.Mount.IP, operate.request.NFS.Mount.Path) // 挂载NFS
	case constants.ActionUnmountNFS:
		return redfish.InspurUmountNFS(operate.ip, operate.username, operate.password) // 卸载NFS
	}
	return fmt.Errorf("浪潮暂不支持操作: %s", operate.request.Action)
}
