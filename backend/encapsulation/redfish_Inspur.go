package encapsulation

import (
	"PeachDRAC/backend/utils"
	"fmt"
	"net/http"
)

type RedfishInspur struct {
	IP       string
	Username string
	Password string
	BaseUrl  string         // 基础地址
	Cookies  []*http.Cookie // 登录后的Cookies
}

// 登录 WEB IPMI 平台(浪潮)
func (r *RedfishInspur) LoginWebIpmi() error {
	url := fmt.Sprintf("https://%s/api/session", r.IP)
	body := fmt.Sprintf("username=%s&password=%s&encrypt_flag=1&login_tag=782504501", r.Username, r.Password)
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"Connection":   "keep-alive",
	}
	_, cookies, err := utils.HttpPostSendAuthRespCookies(url, "", "", body, nil, headers)
	if err != nil {
		return fmt.Errorf("登录失败: %v", err)
	}
	r.Cookies = cookies
	return nil
}

// 下载JNLP文件并且启动
func (r *RedfishInspur) DownloadJnlp() error {
	return nil
}
