package model

import (
	"PeachDRAC/backend/farmework"
	"time"
)

// 密码组配置
type TablePass struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Username  string    `gorm:"not null" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	Port      string    `gorm:"not null" json:"port"`
	Status    bool      `gorm:"not null" json:"status"`   // 是否启用
	Priority  int       `gorm:"not null" json:"priority"` // 优先级,数字越大越高
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
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

func (TablePass) AddOrUpdate(config TablePass) error {
	return farmework.ModuleOrm.Where("id = ?", config.ID).
		Assign(config).FirstOrCreate(&config).Error
}
