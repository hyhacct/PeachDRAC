package model

// 配置响应
type ConfigRespond struct {
	Status bool   `json:"status"` // true-成功 false-失败
	Msg    string `json:"msg"`    // 消息
	Data   any    `json:"data"`   // 数据
}

func Success(data any) ConfigRespond {
	return ConfigRespond{
		Status: true,
		Msg:    "success",
		Data:   data,
	}
}

func Error(msg string) ConfigRespond {
	return ConfigRespond{
		Status: false,
		Msg:    msg,
		Data:   nil,
	}
}
