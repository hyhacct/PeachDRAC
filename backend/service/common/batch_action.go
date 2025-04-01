package common

import (
	"PeachDRAC/backend/constants"
	"PeachDRAC/backend/model"
	"sync"
)

func (s *CommonService) BatchAction(request model.ActionRequest) []model.ActionRespond {

	if len(request.IPs) == 0 {
		return []model.ActionRespond{}
	}

	wg := sync.WaitGroup{}
	wg.Add(len(request.IPs))

	// 结果集
	results := []model.ActionRespond{}

	for _, ip := range request.IPs {
		go func(ip string) {
			defer wg.Done()

			var resultNew model.ActionRespond

			// 获取设备型号
			deviceModel := s.Survey([]string{ip})
			resultNew.IP = ip
			resultNew.Model = deviceModel[0].Model
			resultNew.Action = request.Action
			resultNew.Status = false

			switch deviceModel[0].Model {
			case "戴尔":

				switch request.Action {
				case constants.ActionPowerOn:
					s.DellService.PowerOn(ip, &resultNew) // 执行戴尔的开机操作
				case constants.ActionPowerOff:
					s.DellService.PowerOff(ip, &resultNew) // 执行戴尔的关机操作
				case constants.ActionPowerReset:
					s.DellService.PowerReset(ip, &resultNew) // 执行戴尔的重启操作(冷引导/强制重启)
				case constants.ActionPowerColdBoot:
					s.DellService.PowerReset(ip, &resultNew) // 执行戴尔的重启操作(冷引导/强制重启)
				case constants.ActionPowerForceReset:
					s.DellService.PowerReset(ip, &resultNew) // 执行戴尔的重启操作(冷引导/强制重启)
				}

			case "浪潮":
				// 执行浪潮的操作
			}

			// 将结果添加到结果集
			results = append(results, resultNew)

		}(ip)
	}

	wg.Wait()

	return results
}
