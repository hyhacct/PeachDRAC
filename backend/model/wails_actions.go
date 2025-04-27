package model

type WailsActions struct {
	IPMI         string `json:"ipmi"`
	Status       string `json:"status"` // ready=就绪 success=成功 error=失败
	Model        string `json:"model"`
	Manufacturer string `json:"manufacturer"`
	Sn           string `json:"sn"`
	Action       string `json:"action"`
	Message      string `json:"message"`
}

func WailsActionsSuccess(ipmi string, action string, model string, manufacturer string, sn string) WailsActions {
	return WailsActions{
		IPMI:         ipmi,
		Action:       action,
		Model:        model,
		Manufacturer: manufacturer,
		Sn:           sn,
		Status:       "success",
		Message:      "任务完成",
	}
}

func WailsActionsError(ipmi string, action string, msg string) WailsActions {
	return WailsActions{
		IPMI:         ipmi,
		Action:       action,
		Model:        "",
		Manufacturer: "",
		Sn:           "",
		Status:       "error",
		Message:      msg,
	}
}

func WailsActionsProgress(ipmi string, action string, msg string) WailsActions {
	return WailsActions{
		IPMI:         ipmi,
		Action:       action,
		Model:        "",
		Manufacturer: "",
		Sn:           "",
		Status:       "ready",
		Message:      msg,
	}
}
