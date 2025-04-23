package config

import (
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/utils"
)

func (s *ServiceConfig) AddOrUpdatePass(id int, Username string, Password string, Port string) model.WailsCommunicate {
	if utils.TextIsEmpty(Username) || utils.TextIsEmpty(Password) || utils.TextIsEmpty(Port) {
		return model.WailsError("参数不能为空")
	}

	pass := model.TablePass{
		ID:       id,
		Username: Username,
		Password: Password,
		Port:     Port,
		Status:   true,
	}
	err := (model.TablePass{}).AddOrUpdate(pass)
	if err != nil {
		return model.WailsError("创建密码失败" + err.Error())
	}
	return model.WailsSuccess("创建密码成功", "")
}
