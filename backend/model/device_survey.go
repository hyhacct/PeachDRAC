package model

// 设备探测结果
type DeviceSurvey struct {
	IP           string `json:"ip"`           // IP地址
	ProductName  string `json:"productName"`  // 产品名称
	SerialNumber string `json:"serialNumber"` // 序列号
	Manufacturer string `json:"manufacturer"` // 厂商
	Status       bool   `json:"status"`       // 状态 true:在线 false:离线
}
