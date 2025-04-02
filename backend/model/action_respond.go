package model

type ActionRespond struct {
	IP          string `json:"ip"`          // 设备IP
	ProductName string `json:"productName"` // 设备型号
	Status      bool   `json:"status"`      // 执行的动作
	Action      string `json:"action"`      // 执行的动作
	Result      string `json:"result"`      // 执行的结果
}
