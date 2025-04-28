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

func WailsActionsSuccess(action ModelActions) WailsActions {
	return WailsActions{
		IPMI:         action.IP,
		Action:       action.Action,
		Model:        action.DeviceModel,
		Manufacturer: action.Manufacturer,
		Sn:           action.Sn,
		Status:       "success",
		Message:      "任务完成",
	}
}

func WailsActionsError(action ModelActions, messgae string) WailsActions {
	return WailsActions{
		IPMI:         action.IP,
		Action:       action.Action,
		Model:        action.DeviceModel,
		Manufacturer: action.Manufacturer,
		Sn:           action.Sn,
		Status:       "error",
		Message:      messgae,
	}
}
