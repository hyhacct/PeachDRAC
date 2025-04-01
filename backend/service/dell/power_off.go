package dell

import (
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/utils"
	"fmt"
	"strings"
)

// 关机
func (s *DellService) PowerOff(ip string, result *model.ActionRespond) {

	resp, err := utils.HttpSendGet(fmt.Sprintf("https://%s/data?set=pwState:0", ip))
	if err != nil {
		result.Result = fmt.Sprintf("关机失败: %s", err.Error())
		return
	}

	// 判断响应中是否包含指定字符串
	if strings.Contains(resp, "<status>ok</status>") {
		result.Result = "关机成功"
		result.Status = true // 关机成功
		return
	}

	result.Result = fmt.Sprintf("关机失败: %s", resp)
}
