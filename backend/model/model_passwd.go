package model

import (
	"PeachDRAC/backend/modules"
	"time"
)

type Passwd struct {
	Id          uint      `gorm:"primaryKey;autoIncrement" json:"id"`         // ID
	ProductName string    `gorm:"type:varchar(255)" json:"productName"`       // 设备型号（暂时废弃）
	Username    string    `gorm:"type:varchar(255);not null" json:"username"` // 用户名
	Password    string    `gorm:"type:varchar(255);not null" json:"password"` // 密码
	Port        int       `gorm:"type:int;default:623" json:"port"`           // 端口
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`            // 创建时间
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`            // 更新时间
}

// 表名
func (Passwd) TableName() string {
	return "passwd"
}

// 注册服务
type ConfigPasswd struct {
	orm *modules.Orm
}

func NewConfigPasswd(orm *modules.Orm) *ConfigPasswd {
	return &ConfigPasswd{
		orm: orm,
	}
}

// 添加或更新
func (c *ConfigPasswd) AddOrUpdate(config *Passwd) ConfigRespond {
	err := c.orm.GetDB().Where("id = ?", config.Id).
		Assign(*config).FirstOrCreate(config).Error
	if err != nil {
		return Error("添加或更新失败" + err.Error())
	}
	return Success(nil)
}

// 删除
func (c *ConfigPasswd) Delete(id int) ConfigRespond {
	err := c.orm.GetDB().Where("id = ?", id).Delete(&Passwd{}).Error
	if err != nil {
		return Error("删除失败" + err.Error())
	}
	return Success(nil)
}

// 获取所有
func (c *ConfigPasswd) GetAll() ConfigRespond {
	var configs []Passwd
	if err := c.orm.GetDB().Find(&configs).Error; err != nil {
		return Error("获取所有配置失败" + err.Error())
	}
	return Success(configs)
}
