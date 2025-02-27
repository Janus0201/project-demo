package service

import (
	"context"
	"errors"

	checkout "github.com/MrLittle05/Gomall/app/frontend/hertz_gen/frontend/checkout"
	"github.com/MrLittle05/Gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/MrLittle05/Gomall/app/frontend/utils"
	rpccart "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *checkout.Empty) (resp map[string]any, err error) {
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	if userId == 0 {
		return nil, errors.New("user not login")
	}
	r, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}

	var items []map[string]any
	var total float32
	total = 0.0
	for _, item := range r.Cart.Items {
		p, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: item.ProductId})
		if err != nil {
			continue
		}
		items = append(items, utils.H{
			"Picture":  p.Product.Picture,
			"Name":     p.Product.Name,
			"Price":    p.Product.Price,
			"Quantity": item.Quantity,
		})
		total += p.Product.Price * float32(item.Quantity)
	}
	return utils.H{
		"Title": "Checkout",
		"Items": items,
		"Total": total,
	}, nil
}
