package common

import (
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/service/control"
	"PeachDRAC/backend/utils"
	"sort"
	"sync"
)

// 探测指定IP范围内的设备，并且自动识别型号
func (s *CommonService) Survey(ips []string) []model.DeviceSurvey {

	// 创建一个带缓冲的channel来存储结果
	resultChan := make(chan model.DeviceSurvey, len(ips))
	var wg sync.WaitGroup

	// 检查是否有可用的配置
	

	// 并发处理每个IP
	for _, ip := range ips {
		if !utils.TextIsEmpty(ip) {
			wg.Add(1)
			go func(ipAddr string) {
				defer wg.Done()

				ipAddr = utils.TextTrimSpace(ipAddr) // 去掉空格

				var ipmiControl = control.NewService(ipAddr, "root", "abcd001002", 623)

				if err := ipmiControl.ConnectServer(); err != nil {
					resultChan <- model.DeviceSurvey{
						IP:     ipAddr,
						Status: false,
					}
					return
				}
				system, err := ipmiControl.GetSystem()
				if err != nil {
					resultChan <- model.DeviceSurvey{
						IP:     ipAddr,
						Status: false,
					}
					return
				}
				resultChan <- system

				defer ipmiControl.Close()
			}(ip)
		}
	}

	// 等待所有goroutine完成
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 收集结果
	var respondList []model.DeviceSurvey
	for result := range resultChan {
		respondList = append(respondList, result)
	}

	// 为数组排序
	sort.Slice(respondList, func(i, j int) bool {
		return respondList[i].IP < respondList[j].IP
	})

	return respondList
}
