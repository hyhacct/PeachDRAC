package model

type WailsCommunicate struct {
	Status bool   `json:"Status"` // 状态 true 成功 false 失败
	Msg    string `json:"Msg"`    // 消息
	Data   any    `json:"Data"`   // 数据
}

func WailsSuccess(msg string, data any) WailsCommunicate {
	return WailsCommunicate{
		Status: true,
		Msg:    msg,
		Data:   data,
	}
}

func WailsError(msg string) WailsCommunicate {
	return WailsCommunicate{
		Status: false,
		Msg:    msg,
		Data:   "",
	}
}
