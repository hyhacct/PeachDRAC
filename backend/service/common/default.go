package common

import (
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/service/dell"
	"PeachDRAC/backend/service/inspur"
	"context"
)

type CommonService struct {
	DellService   *dell.DellService
	InspurService *inspur.InspurService
	ConfigPasswd  *model.ConfigPasswd
	ConfigJava    *model.ConfigJava
	Ctx           context.Context
}

func NewService(ctx context.Context, dellService *dell.DellService, inspurService *inspur.InspurService, passwd *model.ConfigPasswd, java *model.ConfigJava) *CommonService {
	return &CommonService{
		Ctx:           ctx,
		DellService:   dellService,
		InspurService: inspurService,
		ConfigPasswd:  passwd,
		ConfigJava:    java,
	}
}
