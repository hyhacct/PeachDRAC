package survey

import "context"

type ServiceSurvey struct {
	ctx context.Context
}

func NewService(ctx context.Context) *ServiceSurvey {
	return &ServiceSurvey{ctx: ctx}
}
