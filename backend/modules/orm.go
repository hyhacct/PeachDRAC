package modules

import (
	"PeachDRAC/backend/constants"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Orm struct {
	db   *gorm.DB
	logc *Logs
}

func NewOrmService(args1 *Logs) *Orm {
	db, err := gorm.Open(sqlite.Open(constants.PathDb))
	if err != nil {
		args1.Error("数据库打开失败: %v", err)
		os.Exit(1)
	}
	return &Orm{
		db:   db,
		logc: args1,
	}
}

func (o *Orm) GetDB() *gorm.DB {
	return o.db
}

func (o *Orm) SyncTable(models ...any) {
	if err := o.db.AutoMigrate(models...); err != nil {
		o.logc.Error("数据库表同步失败: %v", err)
		os.Exit(1)
	}
}
