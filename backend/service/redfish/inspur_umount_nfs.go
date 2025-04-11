package redfish

import (
	"PeachDRAC/backend/utils"
	"fmt"
)

// 浪潮服务器卸载NFS
func InspurUmountNFS(ip string, username string, password string) error {
	token, err := utils.HttpLoginInspur(ip, username, password)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/redfish/v1/Managers/BMC/Vmedia/General", ip)

	headers := map[string]string{
		"Content-Type": "application/json",
		"X-Auth-Token": token,
	}

	body := map[string]interface{}{
		"id":                       1,
		"local_media_support":      0,
		"remote_media_support":     0,
		"same_settings":            0,
		"cd_remote_server_address": "",
		"cd_remote_source_path":    "",
		"cd_remote_share_type":     "",
		"cd_remote_domain_name":    "",
		"cd_remote_user_name":      "",
		"mount_cd":                 0,
		"mount_fd":                 0,
		"fd_remote_server_address": "",
		"fd_remote_source_path":    "",
		"fd_remote_share_type":     "",
		"fd_remote_domain_name":    "",
		"fd_remote_user_name":      "",
		"mount_hd":                 0,
		"hd_remote_server_address": "",
		"hd_remote_source_path":    "",
		"hd_remote_share_type":     "",
		"hd_remote_domain_name":    "",
		"hd_remote_user_name":      "",
	}

	_, err = utils.HttpCommonSend(url, "POST", headers, body, []string{token})
	if err != nil {
		return err
	}
	return nil
}
