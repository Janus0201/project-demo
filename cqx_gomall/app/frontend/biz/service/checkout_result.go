package service

import (
	"context"

	checkout "github.com/MrLittle05/Gomall/app/frontend/hertz_gen/frontend/checkout"
	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutResultService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutResultService(Context context.Context, RequestContext *app.RequestContext) *CheckoutResultService {
	return &CheckoutResultService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutResultService) Run(req *checkout.Empty) (resp map[string]any, err error) {
	return map[string]any{}, nil
}
