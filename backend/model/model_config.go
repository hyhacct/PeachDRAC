package model

import (
	"time"
)

type Config struct {
	Id           uint      `gorm:"primaryKey;default:auto_random()" json:"id"` // ID
	ProductName  string    `gorm:"type:varchar(255)" json:"productName"`       // 产品名称
	Manufacturer string    `gorm:"type:varchar(255)" json:"manufacturer"`      // 制造商
	Username     string    `gorm:"type:varchar(255);not null" json:"username"` // 用户名
	Password     string    `gorm:"type:varchar(255);not null" json:"password"` // 密码
	Port         int       `gorm:"type:int;default:623" json:"port"`           // 端口
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"createdAt"`            // 创建时间
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updatedAt"`            // 更新时间
}

// 表名
func (Config) TableName() string {
	return "config"
}
