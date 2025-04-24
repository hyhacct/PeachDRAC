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
	Port      int       `gorm:"not null" json:"port"`
	Status    bool      `json:"status"`                   // 是否启用
	Priority  int       `gorm:"not null" json:"priority"` // 优先级,数字越大越高
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at" ts_type:"string"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at" ts_type:"string"`
}

func (TablePass) TableName() string {
	return "pass"
}

// 获取所有密码组，默认按照Priority做降序
func (TablePass) GetAll() ([]TablePass, error) {
	var passList []TablePass
	err := farmework.ModuleOrm.Order("priority DESC").Find(&passList).Error
	return passList, err
}

// 获取所有启用的密码组，默认按照Priority做降序
func (TablePass) GetAllEnabled() ([]TablePass, error) {
	var passList []TablePass
	err := farmework.ModuleOrm.Where("status = ?", true).Order("priority DESC").Find(&passList).Error
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

func (TablePass) SwitchStatus(id int, status bool) error {
	return farmework.ModuleOrm.Model(&TablePass{}).Where("id = ?", id).Update("status", status).Error
}
