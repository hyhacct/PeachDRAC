package encapsulation

import (
	"PeachDRAC/backend/farmework"
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"path"
)

type RedfishDell struct {
	IP       string
	Username string
	Password string
	BaseUrl  string         // 基础地址
	Cookies  []*http.Cookie // 登录后的Cookies
	St1      string         // 登录后的ST1
	St2      string         // 登录后的ST2
}

// 取起始接口地址
func (r *RedfishDell) GetBaseUrl() error {
	var (
		url     = fmt.Sprintf("https://%s/redfish/v1/Managers/", r.IP)
		headers = map[string]string{
			"Content-Type": "application/json",
		}
		data model.ModelRedfishDellBaseUrl
	)
	resp, err := utils.HttpPostSendAuth(url, r.Username, r.Password, "", nil, headers)
	if err != nil {
		return fmt.Errorf("获取基础地址失败: %v", err)
	}
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return fmt.Errorf("解析基础地址失败: %v", err)
	}
	if len(data.Members) == 0 {
		return fmt.Errorf("不存在基础地址: %v", err)
	}
	r.BaseUrl = data.Members[0].OdataID
	return nil
}

// 开机
func (r *RedfishDell) PowerOn() error {
	url := fmt.Sprintf("https://%s/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset", r.IP)
	body := `{"ResetType":"On"}`
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	_, err := utils.HttpPostSendAuth(url, r.Username, r.Password, body, nil, headers)
	return err
}

// 优雅关机
func (r *RedfishDell) GracefulShutdown() error {
	url := fmt.Sprintf("https://%s/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset", r.IP)
	body := `{"ResetType":"GracefulShutdown"}`
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	_, err := utils.HttpPostSendAuth(url, r.Username, r.Password, body, nil, headers)
	return err
}

// 强制重启
func (r *RedfishDell) ForceRestart() error {
	url := fmt.Sprintf("https://%s/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset", r.IP)
	body := `{"ResetType":"ForceRestart"}`
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	_, err := utils.HttpPostSendAuth(url, r.Username, r.Password, body, nil, headers)
	return err
}

// 强制关机
func (r *RedfishDell) ForceOff() error {
	url := fmt.Sprintf("https://%s/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset", r.IP)
	body := `{"ResetType":"ForceOff"}`
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	_, err := utils.HttpPostSendAuth(url, r.Username, r.Password, body, nil, headers)
	return err
}

// 挂载NFS
func (r *RedfishDell) MountNFS(nfs string) error {
	url := fmt.Sprintf("https://%s/redfish/v1/Managers/Manager.Embedded.1/Actions/Manager.MountNFS", r.IP)
	body := fmt.Sprintf(`{ "Image": "%s", "Inserted": true }`, nfs)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	_, err := utils.HttpPostSendAuth(url, r.Username, r.Password, body, nil, headers)
	return err
}

// 卸载NFS
func (r *RedfishDell) UnmountNFS() error {
	url := fmt.Sprintf("https://%s/redfish/v1/Managers/Manager.Embedded.1/Actions/Manager.MountNFS", r.IP)
	body := `{}`
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	_, err := utils.HttpPostSendAuth(url, r.Username, r.Password, body, nil, headers)
	return err
}

// 登录 WEB IPMI 平台(R730)
func (r *RedfishDell) LoginWebIpmiR730() error {
	url := fmt.Sprintf("https://%s/data/login", r.IP)
	body := fmt.Sprintf("user=%s&password=%s", r.Username, r.Password)
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	resp, cookies, err := utils.HttpPostSendAuthRespCookies(url, r.Username, r.Password, body, nil, headers)
	if err != nil {
		return fmt.Errorf("登录失败: %v", err)
	}
	farmework.ModuleLogs.Info("登录响应", resp)
	r.St1, r.St2, err = utils.TextGetSt1St2(resp)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	farmework.ModuleLogs.Info("登录成功", r.St1, r.St2)
	r.Cookies = cookies
	return nil
}

// 下载JNLP文件并且启动
func (r *RedfishDell) DownloadJnlp() error {
	url := fmt.Sprintf("https://%s/viewer.jnlp(%s@0@%s,abcd,@1234567890@ST1=%s)", r.IP, r.IP, r.IP, r.St1)
	resp, err := utils.HttpPostSendAuth(url, r.Username, r.Password, "", r.Cookies, nil)
	if err != nil {
		return fmt.Errorf("下载JNLP文件失败: %v", err)
	}
	// 取JNLP文件目录
	dir_path, err := utils.MkdirJnlp()
	if err != nil {
		return fmt.Errorf("创建JNLP文件目录失败: %v", err)
	}

	// 将响应内容写入到目录下
	fileName := path.Join(dir_path, fmt.Sprintf("%s.jnlp", utils.RandomSuid()))
	err = utils.FileWriteToPath(fileName, []byte(resp))
	if err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}

	// 执行下载的JNLP文件
	err = exec.Command("javaws", fileName).Run()
	if err != nil {
		return fmt.Errorf("执行JNLP文件失败: %v", err)
	}
	return nil
}
