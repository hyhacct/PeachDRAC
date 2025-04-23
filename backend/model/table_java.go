package model

import (
	"PeachDRAC/backend/farmework"
	"time"
)

type TableJava struct {
	ID        int       `gorm:"primary_key" json:"id"`   // 主键
	Path      string    `gorm:"not null" json:"path"`    // 路径
	Version   string    `gorm:"not null" json:"version"` // 版本号
	Allot     string    `gorm:"not null" json:"allot"`   // 分配给某个厂商的服务器去使用
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (TableJava) TableName() string {
	return "java"
}

func (TableJava) GetAll() ([]TableJava, error) {
	var javaList []TableJava
	err := farmework.ModuleOrm.Find(&javaList).Error
	return javaList, err
}

func (TableJava) GetByID(id int) (TableJava, error) {
	var java TableJava
	err := farmework.ModuleOrm.Where("id = ?", id).First(&java).Error
	return java, err
}

func (TableJava) Delete(id int) error {
	return farmework.ModuleOrm.Where("id = ?", id).Delete(&TableJava{}).Error
}

func (TableJava) AddOrUpdate(config TableJava) error {
	return farmework.ModuleOrm.Where("id = ?", config.ID).
		Assign(config).FirstOrCreate(&config).Error
}
