package orm

import (
	"PeachDRAC/backend/constants"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Sqlite *SQLite

type SQLite struct {
	err error
	Raw *gorm.DB
}

func NewSQLite() *SQLite {
	t := &SQLite{}

	t.Raw, t.err = gorm.Open(sqlite.Open(constants.PathDb), &gorm.Config{})
	if t.err != nil {
		log.Printf("failed to connect database: %v", t.err)
		os.Exit(1)
	}
	return t
}

func (t *SQLite) Close() {
	if t.Raw != nil {
		sqlDB, err := t.Raw.DB()
		if err != nil {
			log.Printf("failed to get sql.DB: %v", err)
			return
		}
		sqlDB.Close()
	}
}

func (t *SQLite) SyncTable(models ...any) error {
	if t.Raw == nil {
		return fmt.Errorf("database connection is not established")
	}

	if err := t.Raw.AutoMigrate(models...); err != nil {
		return err
	}

	return nil
}
