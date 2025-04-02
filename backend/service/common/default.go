package common

import (
	"PeachDRAC/backend/orm"
	"PeachDRAC/backend/service/dell"
	"PeachDRAC/backend/service/inspur"
	"context"
)

type CommonService struct {
	DellService   *dell.DellService
	InspurService *inspur.InspurService
	OrmService    *orm.SQLite
	Ctx           context.Context
}

func NewService(ctx context.Context, dellService *dell.DellService, inspurService *inspur.InspurService, ormService *orm.SQLite) *CommonService {
	return &CommonService{
		Ctx:           ctx,
		DellService:   dellService,
		InspurService: inspurService,
		OrmService:    ormService,
	}
}
