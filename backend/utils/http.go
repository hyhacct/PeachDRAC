package utils

import (
	"crypto/tls"
	"errors"

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

// 发起通用请求，响应结果以string输出
func HttpCommonSend(url string, method string, headers map[string]string, body interface{}, auth []string) (string, error) {
	var (
		resp *resty.Response
		err  error
	)
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) // 忽略ssl
	request := client.R().SetBody(body).SetHeaders(headers)

	// 如果auth不为空，则设置basic auth
	if len(auth) == 2 {
		request.SetBasicAuth(auth[0], auth[1])
	}

	// 链式调用
	switch method {
	case "GET":
		resp, err = request.Get(url)
	case "POST":
		resp, err = request.Post(url)
	case "PUT":
		resp, err = request.Put(url)
	case "DELETE":
		resp, err = request.Delete(url)
	}
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

// 登录浪潮Redfish获取临时Token
func HttpLoginInspur(url string, username string, password string) (string, error) {
	var request struct {
		UserName string `json:"UserName"`
		Password string `json:"Password"`
	}
	request.UserName = username
	request.Password = password

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) // 忽略ssl

	resp, err := client.R().SetBody(request).SetHeaders(headers).Post(url)
	if err != nil {
		return "", err
	}

	// 从响应头中获取Token
	token := resp.Header().Get("X-Auth-Token")
	if TextIsEmpty(token) {
		return "", errors.New("获取Token失败")
	}
	return token, nil
}
