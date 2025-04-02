package config

import (
	"PeachDRAC/backend/orm"
)

type ConfigService struct {
	orm *orm.SQLite
}

func NewConfigService(orm *orm.SQLite) *ConfigService {
	return &ConfigService{orm: orm}
}
