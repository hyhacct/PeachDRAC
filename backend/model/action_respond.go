package model

type ActionRespond struct {
	IP     string `json:"ip"`
	Model  string `json:"model"`
	Status bool   `json:"status"`
	Action string `json:"action"`
	Result string `json:"result"`
}
