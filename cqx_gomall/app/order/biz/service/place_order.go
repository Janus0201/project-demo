package service

import (
	"context"
	"fmt"

	"github.com/MrLittle05/Gomall/app/order/biz/dal/mysql"
	"github.com/MrLittle05/Gomall/app/order/biz/model"
	order "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	if len(req.OrderItems) == 0 {
		err = fmt.Errorf("OrderItems empty")
		return
	}

	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		orderId, _ := uuid.NewUUID()

		o := &model.Order{
			OrderId:      orderId.String(),
			OrderState:   model.OrderStatePlaced,
			UserId:       req.UserId,
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}
		if req.Address != nil {
			a := req.Address
			o.Consignee.Country = a.Country
			o.Consignee.State = a.State
			o.Consignee.City = a.City
			o.Consignee.StreetAddress = a.StreetAddress
		}
		if err := tx.Create(o).Error; err != nil {
			klog.Errorf("Failed to create order for user %d: %v", req.UserId, err)
			return err
		}

		var itemList []*model.OrderItem
		for _, v := range req.OrderItems {
			itemList = append(itemList, &model.OrderItem{
				OrderIdRefer: o.OrderId,
				ProductId:    v.Item.ProductId,
				Quantity:     int32(v.Item.Quantity),
				Cost:         v.Cost,
			})
		}
		if err := tx.Create(&itemList).Error; err != nil {
			klog.Errorf("Failed to create order items for user %d: %v", req.UserId, err)
			return err
		}
		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: orderId.String(),
			},
		}

		return nil
	})
	klog.Infof("place order success for user %d", req.UserId)
	return
}
