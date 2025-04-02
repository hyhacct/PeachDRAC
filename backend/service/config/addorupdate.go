package config

import (
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/orm"
)

func (c *ConfigService) AddOrUpdate(config *model.Config) model.ConfigRespond {
	err := orm.Sqlite.Raw.Where("id = ?", config.Id).
		Assign(*config).FirstOrCreate(config).Error
	if err != nil {
		return model.Error("添加或更新失败" + err.Error())
	}
	return model.Success(nil)
}
