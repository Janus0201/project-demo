package main

import (
	"context"

	"github.com/MrLittle05/Gomall/app/payment/biz/service"
	payment "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/klog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	resp, err = service.NewChargeService(ctx).Run(req)
	if err != nil {
		klog.Error(err)
		return nil, status.Errorf(codes.InvalidArgument, "business error: %v", err)
	}
	return resp, err
}
