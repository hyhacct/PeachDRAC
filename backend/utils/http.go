package utils

import (
	"crypto/tls"

	"github.com/go-resty/resty/v2"
)

// 发起http请求
func HttpSendGet(url string) (string, error) {
	client := resty.New()

	// 忽略ssl
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	resp, err := client.R().Get(url)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}
