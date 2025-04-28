package interfaces

import (
	"PeachDRAC/backend/utils"
	"fmt"
)

// 挂载NFS
func (s *InterfacesDefault) DellJobMountNfs(nfs string) error {
	url := fmt.Sprintf("https://%s/redfish/v1/Managers/iDRAC.Embedded.1/VirtualMedia/CD/Actions/VirtualMedia.InsertMedia", s.Address)
	body := fmt.Sprintf(`{"id": "CD", "Image": "%s", "Inserted": true }`, nfs)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	_, err := utils.HttpPostSendAuth(url, s.Username, s.Password, body, nil, headers)
	return err
}
