package utils

import (
	"fmt"
	"strings"
)

// 是否浪潮服务器
func IdracIsInspur(ip string) bool {
	html, err := HttpSendGet(fmt.Sprintf("https://%s/", ip))
	if err != nil {
		return false
	}
	return strings.Contains(html, "inspur_logo") // 带有inspur_logo字样的是浪潮服务器
}

// 是否戴尔服务器
func IdracIsDell(ip string) bool {
	html, err := HttpSendGet(fmt.Sprintf("https://%s/locale/locale_zh.json", ip))
	if err != nil {
		return false
	}
	return strings.Contains(html, "support.dell.com") // 带有support.dell.com官网的是戴尔服务器
}
