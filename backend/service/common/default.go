package common

import (
	"PeachDRAC/backend/service/dell"
	"PeachDRAC/backend/service/inspur"
	"context"
)

type CommonService struct {
	DellService   *dell.DellService
	InspurService *inspur.InspurService
	Ctx           context.Context
}

func NewService(ctx context.Context, dellService *dell.DellService, inspurService *inspur.InspurService) *CommonService {
	return &CommonService{
		Ctx:           ctx,
		DellService:   dellService,
		InspurService: inspurService,
	}
}
