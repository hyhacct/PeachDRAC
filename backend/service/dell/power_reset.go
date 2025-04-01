package dell

import (
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/utils"
	"fmt"
	"strings"
)

// 重启/冷引导/强制重启
func (s *DellService) PowerReset(ip string, result *model.ActionRespond) {
	resp, err := utils.HttpSendGet(fmt.Sprintf("https://%s/data?set=pwState:2", ip))
	if err != nil {
		result.Result = fmt.Sprintf("重启失败: %s", err.Error())
		return
	}

	// 判断响应中是否包含指定字符串
	if strings.Contains(resp, "<status>ok</status>") {
		result.Result = "重启成功"
		result.Status = true // 重启成功
		return
	}

	result.Result = "重启失败"
}
