package common

import (
	"PeachDRAC/backend/model"
	"context"
)

type CommonService struct {
	ConfigPasswd *model.ConfigPasswd
	ConfigJava   *model.ConfigJava
	Ctx          context.Context
}

func NewService(ctx context.Context, passwd *model.ConfigPasswd, java *model.ConfigJava) *CommonService {
	return &CommonService{
		Ctx:          ctx,
		ConfigPasswd: passwd,
		ConfigJava:   java,
	}
}
