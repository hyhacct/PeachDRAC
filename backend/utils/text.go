package utils

import "strings"

func TextIsEmpty(text string) bool {
	text = strings.ReplaceAll(text, " ", "")
	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\t", "")
	return text == ""
}
