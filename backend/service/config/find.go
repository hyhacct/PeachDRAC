package config

import (
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/orm"
)

func (c *ConfigService) GetAll() model.ConfigRespond {
	var configs []model.Config
	if err := orm.Sqlite.Raw.Find(&configs).Error; err != nil {
		return model.Error("获取所有配置失败" + err.Error())
	}
	return model.Success(configs)
}
