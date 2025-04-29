package actions

import (
	"PeachDRAC/backend/farmework"
	"context"
)

type ServiceActions struct {
	ctx context.Context
	r   *farmework.Runner
}

func NewService(ctx context.Context) *ServiceActions {
	return &ServiceActions{ctx: ctx}
}
