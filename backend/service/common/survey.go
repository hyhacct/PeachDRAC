package common

import (
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/utils"
	"sort"
	"sync"
)

// 探测指定IP范围内的设备，并且自动识别型号
func (s *CommonService) Survey(ips []string) []model.DeviceModel {

	// 创建一个带缓冲的channel来存储结果
	resultChan := make(chan model.DeviceModel, len(ips))
	var wg sync.WaitGroup

	// 并发处理每个IP
	for _, ip := range ips {
		if !utils.TextIsEmpty(ip) {
			wg.Add(1)
			go func(ipAddr string) {
				defer wg.Done()
				// 去掉空格
				ipAddr = utils.TextTrimSpace(ipAddr)

				var deviceModel string
				if utils.IdracIsDell(ipAddr) {
					deviceModel = "戴尔"
				} else if utils.IdracIsInspur(ipAddr) {
					deviceModel = "浪潮"
				} else {
					deviceModel = "未知/离线"
				}

				resultChan <- model.DeviceModel{
					IP:    ipAddr,
					Model: deviceModel,
				}
			}(ip)
		}
	}

	// 等待所有goroutine完成
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 收集结果
	var respondList []model.DeviceModel
	for result := range resultChan {
		respondList = append(respondList, result)
	}

	// 为数组排序
	sort.Slice(respondList, func(i, j int) bool {
		return respondList[i].IP < respondList[j].IP
	})

	return respondList
}
