package model

type ModelActions struct {
	IP           string `json:"ip"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Action       string `json:"action"`
	Fan          int    `json:"fan"`
	Nfs          string `json:"nfs"`
	
	DeviceModel  string `json:"device_model"` // 设备型号,比如 PowerEdge R740xd
	Sn           string `json:"sn"`           // 序列号SN
	Manufacturer string `json:"manufacturer"` // 制造商,比如 DELL
}
