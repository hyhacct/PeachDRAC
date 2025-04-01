package model

type ActionRequest struct {
	Action string   `json:"action"`
	IPs    []string `json:"ips"`
	Fan    struct {
		Speed int `json:"speed"` // 调整风扇的转速，如果为-1则表示自适应
	} `json:"fan"`
	NFS struct {
		Mount struct {
			IP   string `json:"ip"`   // 挂载NFS的IP
			Path string `json:"path"` // 挂载NFS的路径
		} `json:"mount"`
	} `json:"nfs"`
}
