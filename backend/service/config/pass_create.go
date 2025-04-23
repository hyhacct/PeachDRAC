package config

import (
	"PeachDRAC/backend/model"
)

func (s *ServiceConfig) CreatePass(Username string, Password string, Port string) model.WailsCommunicate {
	pass := model.TablePass{
		Username: Username,
		Password: Password,
		Port:     Port,
		Status:   true,
	}
	err := (model.TablePass{}).Create(pass)
	if err != nil {
		return model.WailsError("创建密码失败" + err.Error())
	}
	return model.WailsSuccess("创建密码成功", "")
}
