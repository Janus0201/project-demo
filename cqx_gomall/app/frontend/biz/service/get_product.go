package service

import (
	"context"

	product "github.com/MrLittle05/Gomall/app/frontend/hertz_gen/frontend/product"
	"github.com/MrLittle05/Gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.GetProductReq) (resp map[string]any, err error) {
	p, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	// map[string]interface{} 和 map[string]any 等价
	// 在 Web 开发中，map[string]interface{} 常用于将数据传递给模板引擎
	return utils.H{
		"item": p.Product,
	}, nil
}
