package dell

import (
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/utils"
	"fmt"
	"strings"
)

// 开机
func (s *DellService) PowerOn(ip string, result *model.ActionRespond) {
	resp, err := utils.HttpSendGet(fmt.Sprintf("https://%s/data?set=pwState:1", ip))
	if err != nil {
		result.Result = fmt.Sprintf("开机失败: %s", err.Error())
		return
	}

	// 判断响应中是否包含指定字符串
	if strings.Contains(resp, "<status>ok</status>") {
		result.Result = "开机成功"
		result.Status = true // 开机成功
		return
	}

	result.Result = "开机失败"
}
