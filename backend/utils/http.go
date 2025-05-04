package utils

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// 统一配置
func HttpTlsConfig() *resty.Client {
	client := resty.New()
	// 配置 TLS 设置
	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS12,
		MaxVersion:         tls.VersionTLS13,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		},
	})
	// 设置超时和重试
	client.SetTimeout(30 * time.Second)
	client.SetRetryCount(5).SetRetryWaitTime(5 * time.Second)
	return client
}

// 发送带身份验证的Get请求
func HttpGetSendAuth(url string, username string, password string, body string, cookies []*http.Cookie, headers map[string]string) (string, error) {
	client := HttpTlsConfig()
	resp, err := client.R().SetBody(body).SetCookies(cookies).SetHeaders(headers).SetBasicAuth(username, password).Get(url)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

// 发送带身份验证的Post请求
func HttpPostSendAuth(url string, username string, password string, body string, cookies []*http.Cookie, headers map[string]string) (string, error) {
	client := HttpTlsConfig()
	resp, err := client.R().SetBody(body).SetCookies(cookies).SetHeaders(headers).SetBasicAuth(username, password).Post(url)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

// 发送带身份验证的Post请求
func HttpPostSendAuthRespCookies(url string, username string, password string, body string, cookies []*http.Cookie, headers map[string]string) (string, []*http.Cookie, error) {
	client := HttpTlsConfig()
	resp, err := client.R().SetBody(body).SetCookies(cookies).SetHeaders(headers).SetBasicAuth(username, password).Post(url)
	if err != nil {
		return "", nil, err
	}
	return resp.String(), resp.Cookies(), nil
}
