package redfish

import (
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/utils"
	"encoding/json"
	"fmt"
	"strings"
)

func DellUmountNFS(ip string, username string, password string) error {
	var (
		responError model.RedfishResponse
		url         = fmt.Sprintf("https://%s/redfish/v1/Managers/iDRAC.Embedded.1/VirtualMedia/CD/Actions/VirtualMedia.EjectMedia", ip)
		body        = `{}`
		headers     = map[string]string{
			"Content-Type": "application/json",
		}
	)
	resp, err := utils.HttpCommonSend(url, "POST", headers, body, []string{username, password})
	if err != nil {
		return err
	}
	// 是否带有error关键字
	if strings.Contains(resp, "error") {
		err = json.Unmarshal([]byte(resp), &responError)
		if err != nil {
			return err
		}
		return fmt.Errorf("卸载失败: %s", responError.Error.Message)
	}
	return nil
}
