package interfaces

import (
	"PeachDRAC/backend/utils"
	"fmt"
)

// 挂载NFS
func (s *InterfacesDefault) DellJobUnMountNfs() error {
	url := fmt.Sprintf("https://%s/redfish/v1/Managers/iDRAC.Embedded.1/VirtualMedia/CD/Actions/VirtualMedia.EjectMedia", s.Address)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	_, err := utils.HttpPostSendAuth(url, s.Username, s.Password, "", nil, headers)
	return err
}
