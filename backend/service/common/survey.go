package common

import (
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/service/control"
	"fmt"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// 探测指定IP范围内的设备，并且自动识别型号
func (s *CommonService) Survey(ips []string) {
	if len(ips) == 0 {
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(len(ips))

	// 获取可用的密码组
	passwdGroups := s.ConfigPasswd.GetAll()

	for _, ip := range ips {
		go func(ip string) {
			defer wg.Done()

			if !passwdGroups.Status {
				runtime.EventsEmit(s.Ctx, "actions", model.ActionRespond{
					IP:     ip,
					Status: false,
					Action: "获取密码组",
					Result: fmt.Sprintf("获取密码组失败: %s", passwdGroups.Msg),
				})
				return
			}
			if len(passwdGroups.Data.([]model.Passwd)) == 0 {
				runtime.EventsEmit(s.Ctx, "actions", model.ActionRespond{
					IP:     ip,
					Status: false,
					Action: "获取密码组",
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
					Action: "获取密码组",
					Result: fmt.Sprintf("已尝试%d个密码组,均登录失败", len(passwdGroups.Data.([]model.Passwd))),
				})
				return
			}

			// 获取设备型号
			_, err := ipmiControl.GetSystem()
			if err != nil {
				runtime.EventsEmit(s.Ctx, "actions", model.ActionRespond{
					IP:     ip,
					Status: false,
					Action: "获取设备型号",
					Result: fmt.Sprintf("获取设备型号失败: %s", err.Error()),
				})
				return
			}

			defer ipmiControl.Close()

		}(ip)
	}
	wg.Wait()
}
