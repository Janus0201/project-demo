package service

import (
	"context"
	"fmt"

	"github.com/MrLittle05/Gomall/app/order/biz/dal/mysql"
	"github.com/MrLittle05/Gomall/app/order/biz/model"
	order "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// Finish your business logic.
	if mysql.DB == nil {
		klog.Errorf("mysql.DB is not initialized")
		return nil, fmt.Errorf("database connection is not established")
	}
	klog.Info("mark order paid success for user %d, order %s", req.UserId, req.OrderId)
	err = model.UpdateOrderState(mysql.DB, s.ctx, req.UserId, req.OrderId, model.OrderStatePaid)
	return
}
