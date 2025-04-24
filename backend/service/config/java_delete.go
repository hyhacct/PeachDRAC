package config

import (
	"PeachDRAC/backend/model"
)

func (s *ServiceConfig) DeleteJava(id int) model.WailsCommunicate {
	err := (model.TableJava{}).Delete(id)
	if err != nil {
		return model.WailsError("删除配置失败" + err.Error())
	}
	return model.WailsSuccess("删除配置成功", "")
}
