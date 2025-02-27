package service

import (
	"context"

	cart "github.com/MrLittle05/Gomall/app/frontend/hertz_gen/frontend/cart"
	"github.com/MrLittle05/Gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/MrLittle05/Gomall/app/frontend/utils"
	rpccart "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddItemService(Context context.Context, RequestContext *app.RequestContext) *AddItemService {
	return &AddItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddItemService) Run(req *cart.AddItemReq) (resp *cart.Empty, err error) {
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
		Item: &rpccart.Item{
			ProductId: req.ProductId,
			Quantity:  req.Quantity,
		},
	})
	if err != nil {
		return nil, err
	}
	return &cart.Empty{}, nil
}
