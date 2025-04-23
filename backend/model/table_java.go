package model

import "PeachDRAC/backend/farmework"

type TableJava struct {
	ID      int    `gorm:"primary_key"` // 主键
	Path    string `gorm:"not null"`    // 路径
	Version string `gorm:"not null"`    // 版本号
	Allot   string `gorm:"not null"`    // 分配给某个厂商的服务器去使用
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

func (TableJava) Update(id int, java TableJava) error {
	return farmework.ModuleOrm.Where("id = ?", id).Updates(&java).Error
}

func (TableJava) Create(java TableJava) error {
	return farmework.ModuleOrm.Create(&java).Error
}
