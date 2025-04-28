package interfaces

import (
	"PeachDRAC/backend/utils"
	"fmt"
)

// 开机
func (s *InterfacesDefault) DellJobPowerOn() error {
	url := fmt.Sprintf("https://%s/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset", s.Address)
	body := `{"ResetType":"On"}`
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	_, err := utils.HttpPostSendAuth(url, s.Username, s.Password, body, s.Cookies, headers)
	return err
}
