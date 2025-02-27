package main

import (
	"context"

	"github.com/MrLittle05/Gomall/app/checkout/biz/service"
	checkout "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/checkout"
	"github.com/cloudwego/kitex/pkg/klog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CheckoutServiceImpl implements the last service interface defined in the IDL.
type CheckoutServiceImpl struct{}

// Checkout implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) Checkout(ctx context.Context, req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	resp, err = service.NewCheckoutService(ctx).Run(req)
	if err != nil {
		klog.Error(err)
		return nil, status.Errorf(codes.InvalidArgument, "business error: %v", err)
	}
	return resp, err
}
