package service

import (
	"context"

	"github.com/MrLittle05/Gomall/app/order/biz/dal/mysql"
	"github.com/MrLittle05/Gomall/app/order/biz/model"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/cart"
	order "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	orders, err := model.ListOrder(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		klog.Errorf("model.ListOrder.err:%v", err)
		return nil, err
	}
	var list []*order.Order
	for _, v := range orders {
		var items []*order.OrderItem
		for _, v := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Cost: v.Cost,
				Item: &cart.Item{
					ProductId: v.ProductId,
					Quantity:  uint32(v.Quantity),
				},
			})
		}
		o := &order.Order{
			OrderId:      v.OrderId,
			UserId:       v.UserId,
			UserCurrency: v.UserCurrency,
			Email:        v.Consignee.Email,
			CreatedAt:    int32(v.CreatedAt.Unix()),
			Address: &order.Address{
				Country:       v.Consignee.Country,
				City:          v.Consignee.City,
				StreetAddress: v.Consignee.StreetAddress,
				ZipCode:       v.Consignee.ZipCode,
			},
			OrderItems: items,
		}
		list = append(list, o)
	}
	resp = &order.ListOrderResp{
		Orders: list,
	}
	return
}
