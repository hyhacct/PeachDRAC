package common

import "PeachDRAC/backend/utils"

// 探测指定IP范围内的设备，并且自动识别型号
func (s *CommonService) Survey(ips []string) interface{} {

	type respond struct {
		IP    string // IP地址
		Model string // 型号
	}

	// 响应列表
	var respondList = make([]respond, 0)

	for _, ip := range ips {
		if !utils.TextIsEmpty(ip) {

			// 去掉空格
			ip = utils.TextTrimSpace(ip)

			if utils.IdracIsDell(ip) {
				respondList = append(respondList, respond{
					IP:    ip,
					Model: "戴尔",
				})
				continue
			}

			if utils.IdracIsInspur(ip) {
				respondList = append(respondList, respond{
					IP:    ip,
					Model: "浪潮",
				})
				continue
			}

			respondList = append(respondList, respond{
				IP:    ip,
				Model: "未知/离线",
			})
		}
	}

	return respondList
}
