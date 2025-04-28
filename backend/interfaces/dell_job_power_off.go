package interfaces

import (
	"PeachDRAC/backend/utils"
	"fmt"
)

// 关机
func (s *InterfacesDefault) DellJobPowerOff() error {
	url := fmt.Sprintf("https://%s/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset", s.Address)
	body := `{"ResetType":"ForceOff"}`
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	_, err := utils.HttpPostSendAuth(url, s.Username, s.Password, body, s.Cookies, headers)
	return err
}
