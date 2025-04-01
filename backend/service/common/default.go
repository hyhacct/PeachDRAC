package common

import (
	"PeachDRAC/backend/service/dell"
	"PeachDRAC/backend/service/inspur"
)

type CommonService struct {
	DellService   *dell.DellService
	InspurService *inspur.InspurService
}

func NewService(dellService *dell.DellService, inspurService *inspur.InspurService) *CommonService {
	return &CommonService{
		DellService:   dellService,
		InspurService: inspurService,
	}
}
