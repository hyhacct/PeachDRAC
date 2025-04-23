package config

import "PeachDRAC/backend/model"

func (s *ServiceConfig) GetAllPass() model.WailsCommunicate {
	list, err := (model.TablePass{}).GetAll()
	if err != nil {
		return model.WailsError(err.Error())
	}
	return model.WailsSuccess("获取密码列表成功", list)
}
