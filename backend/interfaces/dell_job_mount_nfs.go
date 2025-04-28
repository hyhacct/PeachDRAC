package interfaces

import (
	"PeachDRAC/backend/utils"
	"fmt"
)

// 挂载NFS
func (s *InterfacesDefault) DellJobMountNfs(nfs string) error {
	url := fmt.Sprintf("https://%s/redfish/v1/Managers/Manager.Embedded.1/Actions/Manager.MountNFS", s.Address)
	body := fmt.Sprintf(`{ "Image": "%s", "Inserted": true }`, nfs)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	_, err := utils.HttpPostSendAuth(url, s.Username, s.Password, body, nil, headers)
	return err
}
