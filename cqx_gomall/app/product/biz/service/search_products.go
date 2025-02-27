package service

import (
	"context"

	"github.com/MrLittle05/Gomall/app/product/biz/dal/mysql"
	"github.com/MrLittle05/Gomall/app/product/biz/model"
	product "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	p, err := model.NewProductQuery(s.ctx, mysql.DB).SearchProducts(req.Query)
	if err != nil {
		return nil, err
	}
	resp = &product.SearchProductsResp{}
	for _, v := range p {
		resp.Results = append(resp.Results, &product.Product{
			Id:          uint32(v.ID),
			Picture:     string(v.Picture),
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
		})
	}
	return resp, err
}
