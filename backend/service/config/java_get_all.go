package config

import "PeachDRAC/backend/model"

func (s *ServiceConfig) GetAllJava() model.WailsCommunicate {
	list, err := (model.TableJava{}).GetAll()
	if err != nil {
		return model.WailsError(err.Error())
	}
	return model.WailsSuccess("获取配置列表成功", list)
}
