package model

type WailsCommunicate struct {
	Status bool   `json:"status"` // 状态 true 成功 false 失败
	Msg    string `json:"msg"`    // 消息
	Data   string `json:"data"`   // 数据
}

func WailsSuccess(msg string, data string) WailsCommunicate {
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
