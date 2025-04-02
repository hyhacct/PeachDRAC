package config

import (
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/orm"
)

func (c *ConfigService) Delete(id int) model.ConfigRespond {
	err := orm.Sqlite.Raw.Where("id = ?", id).Delete(&model.Config{}).Error
	if err != nil {
		return model.Error("删除失败" + err.Error())
	}
	return model.Success(nil)
}
