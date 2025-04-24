package config

import (
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/utils"
)

func (s *ServiceConfig) AddOrUpdateJava(form model.TableJava) model.WailsCommunicate {
	if utils.TextIsEmpty(form.Title) || utils.TextIsEmpty(form.Path) {
		return model.WailsError("参数不能为空")
	}

	err := (model.TableJava{}).AddOrUpdate(form)
	if err != nil {
		return model.WailsError("创建配置失败" + err.Error())
	}
	return model.WailsSuccess("创建配置成功", "")
}
