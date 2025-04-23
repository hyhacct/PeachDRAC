package model

import (
	"PeachDRAC/backend/farmework"
	"time"
)

// 密码组配置
type TablePass struct {
	ID        int       `gorm:"primary_key"`
	Username  string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	Port      string    `gorm:"not null"`
	Status    bool      `gorm:"not null"` // 是否启用
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (TablePass) TableName() string {
	return "pass"
}

func (TablePass) GetAll() ([]TablePass, error) {
	var passList []TablePass
	err := farmework.ModuleOrm.Find(&passList).Error
	return passList, err
}

func (TablePass) GetByID(id int) (TablePass, error) {
	var pass TablePass
	err := farmework.ModuleOrm.Where("id = ?", id).First(&pass).Error
	return pass, err
}

func (TablePass) Delete(id int) error {
	return farmework.ModuleOrm.Where("id = ?", id).Delete(&TablePass{}).Error
}

func (TablePass) Update(id int, pass TablePass) error {
	return farmework.ModuleOrm.Where("id = ?", id).Updates(&pass).Error
}

func (TablePass) Create(pass TablePass) error {
	return farmework.ModuleOrm.Create(&pass).Error
}
