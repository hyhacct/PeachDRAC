package model

import (
	"PeachDRAC/backend/modules"
	"time"
)

type Java struct {
	Id           uint      `gorm:"primaryKey;autoIncrement" json:"id"`    // ID
	Manufacturer string    `gorm:"type:varchar(255)" json:"manufacturer"` // 制造商
	JavaPath     string    `gorm:"type:varchar(255)" json:"javaPath"`     // Java 路径
	JavaVersion  string    `gorm:"type:varchar(255)" json:"javaVersion"`  // Java 版本
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"createdAt"`       // 创建时间
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updatedAt"`       // 更新时间
}

// 表名
func (Java) TableName() string {
	return "java"
}

// 注册服务
type ConfigJava struct {
	orm *modules.Orm
}

func NewConfigJava(orm *modules.Orm) *ConfigJava {
	return &ConfigJava{
		orm: orm,
	}
}

// 添加或更新
func (c *ConfigJava) AddOrUpdate(config *Java) ConfigRespond {
	err := c.orm.GetDB().Where("manufacturer = ?", config.Manufacturer).
		Assign(*config).FirstOrCreate(config).Error
	if err != nil {
		return Error("添加或更新失败" + err.Error())
	}
	return Success(nil)
}

// 删除
func (c *ConfigJava) Delete(id int) ConfigRespond {
	err := c.orm.GetDB().Where("id = ?", id).Delete(&Java{}).Error
	if err != nil {
		return Error("删除失败" + err.Error())
	}
	return Success(nil)
}

// 获取所有
func (c *ConfigJava) GetAll() ConfigRespond {
	var configs []Java
	if err := c.orm.GetDB().Find(&configs).Error; err != nil {
		return Error("获取所有配置失败" + err.Error())
	}
	return Success(configs)
}
