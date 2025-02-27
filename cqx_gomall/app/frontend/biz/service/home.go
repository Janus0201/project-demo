package service

import (
	"context"

	home "github.com/MrLittle05/Gomall/app/frontend/hertz_gen/frontend/home"
	"github.com/MrLittle05/Gomall/app/frontend/infra/rpc"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (map[string]any, error) {
	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"Title": "Hot Sale",
		"Items": products.Products,
	}, nil
}
