package utils

import (
	"crypto/tls"
	"net/http"

	"github.com/go-resty/resty/v2"
)

// 发送带身份验证的Get请求
func HttpGetSendAuth(url string, username string, password string, body string, cookies []*http.Cookie, headers map[string]string) (string, error) {
	client := resty.New()
	// 忽略ssl证书
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().SetBody(body).SetCookies(cookies).SetHeaders(headers).SetBasicAuth(username, password).Get(url)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

// 发送带身份验证的Post请求
func HttpPostSendAuth(url string, username string, password string, body string, cookies []*http.Cookie, headers map[string]string) (string, error) {
	client := resty.New()
	// 忽略ssl证书
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().SetBody(body).SetCookies(cookies).SetHeaders(headers).SetBasicAuth(username, password).Post(url)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

// 发送带身份验证的Post请求
func HttpPostSendAuthRespCookies(url string, username string, password string, body string, cookies []*http.Cookie, headers map[string]string) (string, []*http.Cookie, error) {
	client := resty.New()
	// 忽略ssl证书
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().SetBody(body).SetCookies(cookies).SetHeaders(headers).SetBasicAuth(username, password).Post(url)
	if err != nil {
		return "", nil, err
	}
	return resp.String(), resp.Cookies(), nil
}
