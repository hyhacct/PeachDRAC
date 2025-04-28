package interfaces

import (
	"PeachDRAC/backend/utils"
	"fmt"
)

// 设置本地 CD 引导项
func (s *InterfacesDefault) DellJobBootSetLocalCd() error {
	url := fmt.Sprintf("https://%s/redfish/v1/Systems/System.Embedded.1", s.Address)
	body := `{ "Boot": {"BootSourceOverrideEnabled": "Once", "BootSourceOverrideTarget": "Cd"}}`
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	_, err := utils.HttpPostSendAuth(url, s.Username, s.Password, body, nil, headers)
	return err
}

// 设置正常引导
func (s *InterfacesDefault) DellJobBootSetNone() error {
	url := fmt.Sprintf("https://%s/redfish/v1/Systems/System.Embedded.1", s.Address)
	body := `{ "Boot": {"BootSourceOverrideEnabled": "None"}}`
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	_, err := utils.HttpPostSendAuth(url, s.Username, s.Password, body, nil, headers)
	return err
}
