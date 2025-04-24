package config

import (
	"PeachDRAC/backend/farmework"
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/utils"
)

func (s *ServiceConfig) AddOrUpdatePass(form model.TablePass) model.WailsCommunicate {
	if utils.TextIsEmpty(form.Username) || utils.TextIsEmpty(form.Password) {
		return model.WailsError("参数不能为空")
	}

	farmework.ModuleLogs.Info("添加密码组", "form = ", form)

	err := (model.TablePass{}).AddOrUpdate(form)
	if err != nil {
		return model.WailsError("创建密码失败" + err.Error())
	}
	return model.WailsSuccess("创建密码成功", "")
}
