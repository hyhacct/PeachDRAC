package farmework

import (
	"PeachDRAC/backend/constants"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	ModuleOrm *gorm.DB
)

func InitOrm() {
	db, err := gorm.Open(sqlite.Open(constants.PathSqlite), &gorm.Config{})
	if err != nil {
		ModuleLogs.Error("数据库初始化失败", err.Error())
		os.Exit(1) // 退出程序
	}
	ModuleOrm = db
	ModuleLogs.Success("数据库初始化成功")
}

func AutoMigrate(tables ...interface{}) {
	err := ModuleOrm.AutoMigrate(tables...)
	if err != nil {
		ModuleLogs.Error("数据库迁移失败", err.Error())
		os.Exit(1) // 退出程序
	}
	ModuleLogs.Success("数据库迁移完成")
}
