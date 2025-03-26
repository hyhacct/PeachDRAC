package modules

import (
	"PeachDRAC/backend/constants"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ModulesOrm struct {
	db          *gorm.DB
	logsService *ModulesLogs
}

func NewModulesOrm(logsService *ModulesLogs) *ModulesOrm {
	return &ModulesOrm{
		logsService: logsService,
	}
}

func (c *ModulesOrm) Init() {
	db, err := gorm.Open(sqlite.Open(constants.PathDb), &gorm.Config{})
	if err != nil {
		c.logsService.Error("初始化本地SQLite数据库失败: %v", err)
		os.Exit(1) // 退出程序
	}
	c.db = db
}

func (c *ModulesOrm) GetDB() *gorm.DB {
	return c.db
}

func (c *ModulesOrm) Close() {
	db, err := c.db.DB()
	if err != nil {
		c.logsService.Error("关闭数据库失败: %v", err)
		os.Exit(1) // 退出程序
	}
	db.Close()
}

// 同步数据库表结构
func (c *ModulesOrm) SyncTables(tables ...interface{}) {
	errStr := c.db.AutoMigrate(tables...).Error
	c.logsService.Info("同步数据库表结构: %v", errStr)
}
