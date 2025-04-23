package config

import "PeachDRAC/backend/model"

func (s *ServiceConfig) UpdatePass(id int, Username string, Password string, Port string) model.WailsCommunicate {
	pass := model.TablePass{
		Username: Username,
		Password: Password,
		Port:     Port,
	}
	err := (model.TablePass{}).Update(id, pass)
	if err != nil {
		return model.WailsError("更新密码失败" + err.Error())
	}
	return model.WailsSuccess("更新密码成功", "")
}
