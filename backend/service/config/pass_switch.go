package config

import (
	"PeachDRAC/backend/farmework"
	"PeachDRAC/backend/model"
	"strconv"
)

func (s *ServiceConfig) SwitchPass(id int, status bool) model.WailsCommunicate {
	err := model.TablePass{}.SwitchStatus(id, status)
	if err != nil {
		return model.WailsError("更新失败" + err.Error())
	}
	farmework.ModuleLogs.Info("更新密码组状态成功", "id = "+strconv.Itoa(id)+", status = "+strconv.FormatBool(status))
	return model.WailsSuccess("更新成功", "")
}
