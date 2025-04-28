package actions

import (
	"PeachDRAC/backend/constants"
	interfaces "PeachDRAC/backend/interfaces"
	"PeachDRAC/backend/model"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (s *ServiceActions) BranchInspur(client *interfaces.InterfacesDefault) {
	var err error

	switch client.Actions.Action {
	// case constants.ActionPowerOn:
	// 	err = client.DellJobPowerOn() // 开机
	default:
		runtime.EventsEmit(s.ctx, constants.EventActions, model.WailsActionsError(client.Actions, "浪潮暂不支持此操作"))
		return
	}
	// 最终检查有没有错误
	if err != nil {
		runtime.EventsEmit(s.ctx, constants.EventActions, model.WailsActionsError(client.Actions, err.Error()))
		return
	}
	runtime.EventsEmit(s.ctx, constants.EventActions, model.WailsActionsSuccess(client.Actions))

}
