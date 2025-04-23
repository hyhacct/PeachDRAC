package config

import (
	"PeachDRAC/backend/model"
)

func (s *ServiceConfig) DeletePass(id int) model.WailsCommunicate {
	err := (model.TablePass{}).Delete(id)
	if err != nil {
		return model.WailsError("删除密码失败" + err.Error())
	}
	return model.WailsSuccess("删除密码成功", "")
}
