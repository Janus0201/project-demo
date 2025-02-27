package service

import (
	"context"

	"github.com/MrLittle05/Gomall/app/product/biz/dal/mysql"
	"github.com/MrLittle05/Gomall/app/product/biz/model"
	product "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	c, err := model.NewCategoryQuery(s.ctx, mysql.DB).GetProductByCategoryName(req.CategoryName)
	if err != nil {
		return nil, err
	}
	resp = &product.ListProductsResp{}
	for _, v1 := range c {
		for _, v2 := range v1.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(v2.ID),
				Picture:     string(v2.Picture),
				Name:        v2.Name,
				Description: v2.Description,
				Price:       v2.Price,
			})
		}
	}
	return resp, err
}
