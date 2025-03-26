package utils

import "strings"

// 去掉字符串的空格
func TextTrimSpace(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

// 判断字符串是否为空
func TextIsEmpty(str string) bool {
	return TextTrimSpace(str) == ""
}
