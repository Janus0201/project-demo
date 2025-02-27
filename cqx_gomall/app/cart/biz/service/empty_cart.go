package service

import (
	"context"

	"github.com/MrLittle05/Gomall/app/cart/biz/dal/mysql"
	"github.com/MrLittle05/Gomall/app/cart/biz/model"
	cart "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	err = model.NewCartQuery(s.ctx, mysql.DB).EmptyCart(req.UserId)
	if err != nil {
		klog.Errorf("Empty cart failed for user %d: %v", req.UserId, err)
		return nil, kerrors.NewBizStatusError(50001, err.Error())
	}
	return &cart.EmptyCartResp{}, nil
}
