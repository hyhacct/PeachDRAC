package actions

import (
	"PeachDRAC/backend/constants"
	interfaces "PeachDRAC/backend/interfaces"
	"PeachDRAC/backend/model"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (s *ServiceActions) BranchDell(client *interfaces.InterfacesDefault) {
	var err error

	switch client.Actions.Action {
	case constants.ActionPowerOn:
		err = client.DellJobPowerOn() // 开机
	case constants.ActionRestart:
		err = client.DellJobPowerRestart() // 重启
	case constants.ActionPowerOff:
		err = client.DellJobPowerOff() // 关机
	case constants.ActionMountNFS:
		err = client.DellJobMountNfs(client.Actions.Nfs) // 挂载NFS
	case constants.ActionUnmountNFS:
		err = client.DellJobUnMountNfs() // 卸载NFS
		if err == nil {
			err = client.DellJobBootSetNone() // 恢复正常引导
		}
	case constants.ActionStartJavaConsole:
		err = client.DellJobLoginWebR730() // 登录Web Ipmi R730
		if err == nil {
			err = client.DellJobDownloadJnlp() // 下载JNLP文件
		}
	case constants.ActionSetFan:
		err = client.DellJobSetFan() // 手动设置风扇转速
	case constants.ActionSetFanAuto:
		err = client.DellJobSetFanAuto() // 风扇自适应
	default:
		runtime.EventsEmit(s.ctx, constants.EventActions, model.WailsActionsError(client.Actions, "戴尔暂不支持此操作"))
		return
	}
	// 最终检查有没有错误
	if err != nil {
		runtime.EventsEmit(s.ctx, constants.EventActions, model.WailsActionsError(client.Actions, err.Error()))
		return
	}
	runtime.EventsEmit(s.ctx, constants.EventActions, model.WailsActionsSuccess(client.Actions))
}
