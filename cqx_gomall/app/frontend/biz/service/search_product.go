package service

import (
	"context"

	product "github.com/MrLittle05/Gomall/app/frontend/hertz_gen/frontend/product"
	"github.com/MrLittle05/Gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type SearchProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductService(Context context.Context, RequestContext *app.RequestContext) *SearchProductService {
	return &SearchProductService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductService) Run(req *product.SearchProductReq) (resp map[string]any, err error) {
	p, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{
		Query: req.Query,
	})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"title": "Search",
		"q":     req.Query,
		"items": p.Results,
	}, nil
}
