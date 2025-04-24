package model

type WailsTask struct {
	ID    string   `json:"id"`    // 任务ID,一般情况下根据IP地址来的
	Done  bool     `json:"done"`  // 任务是否完成
	Exit  bool     `json:"exit"`  // 任务是否退出
	Msg   string   `json:"msg"`   // 消息(如果未完成那就将这个消息更新过去)
	Login bool     `json:"login"` // 是否登录成功?
	Args  []string `json:"args"`  // 附带参数，比如厂商，型号，SN等等，当然也可以不需要
}

// 任务顺利完成
func WailsTaskSuccess(id string, msg string, args []string) WailsTask {
	return WailsTask{
		ID:    id,
		Done:  true,
		Exit:  true,
		Msg:   msg,
		Args:  args,
		Login: true,
	}
}

// 任务异常退出
func WailsTaskExit(loging bool, id string, msg string) WailsTask {
	return WailsTask{
		ID:    id,
		Done:  false,
		Exit:  true,
		Msg:   msg,
		Login: loging,
	}
}

// 任务进行中(仅发送消息)
func WailsTaskProgress(id string, msg string) WailsTask {
	return WailsTask{
		ID:    id,
		Done:  false,
		Exit:  false,
		Msg:   msg,
		Login: true,
	}
}
