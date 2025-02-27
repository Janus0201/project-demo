package service

import (
	"context"

	"github.com/MrLittle05/Gomall/app/cart/biz/dal/mysql"
	"github.com/MrLittle05/Gomall/app/cart/biz/model"
	cart "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	c, err := model.NewCartQuery(s.ctx, mysql.DB).GetCart(req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50002, err.Error())
	}
	return &cart.GetCartResp{Cart: c}, err
}
