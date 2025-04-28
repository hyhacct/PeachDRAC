package interfaces

import (
	"PeachDRAC/backend/farmework"
	"PeachDRAC/backend/utils"
	"fmt"
)

// 登录R730 Web界面
func (s *InterfacesDefault) DellJobLoginWebR730() error {
	url := fmt.Sprintf("https://%s/data/login", s.Address)
	body := fmt.Sprintf("user=%s&password=%s", s.Username, s.Password)
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	resp, cookies, err := utils.HttpPostSendAuthRespCookies(url, s.Username, s.Password, body, nil, headers)
	if err != nil {
		farmework.ModuleLogs.Error("IPMI:", s.Address, "用户名:", s.Username, "密码:", s.Password, "登录失败:", err, "响应:", resp)
		return fmt.Errorf("登录失败: %v", err)
	}
	s.St1, s.St2, err = utils.TextGetSt1St2(resp)
	if err != nil {
		farmework.ModuleLogs.Error("IPMI:", s.Address, "用户名:", s.Username, "密码:", s.Password, "获取ST1和ST2失败:", err, "响应:", resp)
		return fmt.Errorf("%v", err)
	}
	s.Cookies = cookies
	return nil
}
