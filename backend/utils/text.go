package utils

import (
	"encoding/xml"
	"fmt"
	"strings"
)

func TextIsEmpty(text string) bool {
	text = strings.ReplaceAll(text, " ", "")
	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\t", "")
	return text == ""
}

// <?xml version="1.0" encoding="UTF-8"?> <root> <status>ok</status> <authResult>0</authResult> <forwardUrl>index.html?ST1=5c3d8b7ba9f10fb1af7b99343e2bd18e,ST2=eac8d890865ba139f21d4db2e4d459b4</forwardUrl> </root>

func TextGetSt1St2(text string) (string, string, error) {
	type root struct {
		XMLName    xml.Name `xml:"root"`
		Status     string   `xml:"status"`
		AuthResult int      `xml:"authResult"`
		ForwardURL string   `xml:"forwardUrl"`
	}
	var data root
	err := xml.Unmarshal([]byte(text), &data)
	if err != nil {
		return "", "", fmt.Errorf("解析登录数据失败: %v", err)
	}
	data.ForwardURL = strings.ReplaceAll(data.ForwardURL, "index.html?", "")
	urlParts := strings.Split(data.ForwardURL, ",")
	var st1, st2 string
	for _, part := range urlParts {
		if strings.HasPrefix(part, "ST1=") {
			st1 = strings.TrimPrefix(part, "ST1=")
		} else if strings.HasPrefix(part, "ST2=") {
			st2 = strings.TrimPrefix(part, "ST2=")
		}
	}
	if TextIsEmpty(st1) || TextIsEmpty(st2) {
		return "", "", fmt.Errorf("ST1或ST2为空")
	}
	return st1, st2, nil
}
