package interfaces

import (
	"PeachDRAC/backend/utils"
	"fmt"
)

// 重启
func (s *InterfacesDefault) DellJobPowerRestart() error {
	url := fmt.Sprintf("https://%s/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset", s.Address)
	body := `{"ResetType":"ForceRestart"}`
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	_, err := utils.HttpPostSendAuth(url, s.Username, s.Password, body, s.Cookies, headers)
	return err
}
