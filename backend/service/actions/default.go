package actions

import "context"

type ServiceActions struct {
	ctx context.Context
}

func NewService(ctx context.Context) *ServiceActions {
	return &ServiceActions{ctx: ctx}
}
