package survey

import (
	"PeachDRAC/backend/constants"
	"PeachDRAC/backend/farmework"
	"PeachDRAC/backend/model"
	"fmt"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

/*
开始探测
*/
func (s *ServiceSurvey) StartSurvey(ips []string) model.WailsCommunicate {

	// 得到所有密码组
	pass_list, err := model.TablePass{}.GetAllEnabled()
	if err != nil {
		return model.WailsError("获取密码组失败")
	}

	if len(pass_list) == 0 {
		return model.WailsError("没有密码组")
	}

	// 循环探测IP地址，一般地址都有二百多个，需要做多线程并发
	wg := sync.WaitGroup{}
	for _, ip := range ips {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()

			var (
				client    = &farmework.IPMI{} // IPMI客户端
				isSuccess = false             // 是否成功
			)

			for _, item := range pass_list {
				if err := client.Connect(ip, item.Username, item.Password, item.Port); err != nil {
					continue
				}
				isSuccess = true
				break
			}

			if !isSuccess {
				runtime.EventsEmit(s.ctx, constants.EventTask, model.WailsTaskExit(false, ip, fmt.Sprintf("已尝试%d个密码组,均失败", len(pass_list))))
				return
			}

			// 读取硬件信息
			deviceModel, sn, manufacturer, err := client.GetModelAndSN()
			if err != nil {
				runtime.EventsEmit(s.ctx, constants.EventTask, model.WailsTaskExit(true, ip, fmt.Sprintf("读取硬件信息失败: %s", err)))
				return
			}
			args := []string{deviceModel, sn, manufacturer}
			// model.WailsTaskSuccess(ip, fmt.Sprintf("读取硬件信息成功"), args)
			runtime.EventsEmit(s.ctx, constants.EventTask, model.WailsTaskSuccess(ip, "任务结束,已完成", args))
		}(ip)
	}

	// 等待所有探测完成
	wg.Wait()

	return model.WailsSuccess("探测完成", "")
}
