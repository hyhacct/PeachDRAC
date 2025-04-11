package redfish

import (
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/utils"
	"encoding/json"
	"fmt"
	"strings"
)

func DellMountNFS(ip string, username string, password string, image string) error {
	var (
		responError model.RedfishResponse
		url         = fmt.Sprintf("https://%s/redfish/v1/Managers/iDRAC.Embedded.1/VirtualMedia/CD/Actions/VirtualMedia.InsertMedia", ip)
		body        = fmt.Sprintf(`{"Image": "%s", "Inserted": true}`, image)
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
		return fmt.Errorf("挂载失败: %s", responError.Error.Message)
	}
	return nil
}
