package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/MrLittle05/Gomall/app/checkout/rpc"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/checkout"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/order"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/payment"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"google.golang.org/grpc/status"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Get item list from cart
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})

	if err != nil {
		return nil, kerrors.NewBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Cart.Items == nil {
		return nil, kerrors.NewBizStatusError(5004001, "Cart is empty")
	}

	// Calculate total price
	var (
		oi    []*order.OrderItem
		total float32
	)
	for _, cartItem := range cartResult.Cart.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: cartItem.ProductId})
		if resultErr != nil {
			klog.Error(resultErr)
			err = resultErr
			return
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		cost := p.Price * float32(cartItem.Quantity)
		total += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.Item{ProductId: cartItem.ProductId, Quantity: cartItem.Quantity},
			Cost: cost,
		})
	}

	// create order
	orderReq := &order.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: "USD",
		OrderItems:   oi,
		Email:        req.Phone,
	}
	if req.Address != nil {
		addr := req.Address
		zipCodeInt, _ := strconv.Atoi(addr.ZipCode)
		orderReq.Address = &order.Address{
			StreetAddress: addr.StreetAddress,
			City:          addr.City,
			Country:       addr.Country,
			State:         addr.State,
			ZipCode:       int32(zipCodeInt),
		}
	}
	orderResult, err := rpc.OrderClient.PlaceOrder(s.ctx, orderReq)
	if err != nil {
		err = fmt.Errorf("PlaceOrder.err: %v", err)
		return
	}
	klog.Info("orderResult: ", orderResult)

	// Payment
	var orderId string
	if orderResult != nil || orderResult.Order != nil {
		orderId = orderResult.Order.OrderId
	}
	payResult, err := rpc.PayClient.Charge(s.ctx, &payment.ChargeReq{
		Amount:     total,
		UserId:     req.UserId,
		OrderId:    orderId,
		CreditCard: req.CreditCard,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			klog.Errorf("Payment failed for user %d, order %s: code=%s, message=%s", req.UserId, orderId, st.Code(), st.Message())
		} else {
			klog.Errorf("Payment failed for user %d, order %s: %v", req.UserId, orderId, err)
		}
		return nil, kerrors.NewBizStatusError(5005003, err.Error())
	}
	if payResult == nil || payResult.TransactionId == "" {
		klog.Errorf("Payment result not found for user %d, order %s", req.UserId, orderId)
		return nil, kerrors.NewBizStatusError(5004003, "Payment failed")
	}
	klog.Infof("Payment success for user %d, order %s", req.UserId, orderId)

	// Update order state
	_, err = rpc.OrderClient.MarkOrderPaid(s.ctx, &order.MarkOrderPaidReq{
		UserId:  req.UserId,
		OrderId: orderId,
	})
	if err != nil {
		klog.Errorf("Update order state failed for user %d, order %s: %v", req.UserId, orderId, err)
		return nil, kerrors.NewBizStatusError(5005002, err.Error())
	}
	klog.Infof("Update order state success for user %d, order %s", req.UserId, orderId)

	// Empty cart
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Errorf("Empty cart failed for user %d: %v", req.UserId, err)
		return nil, kerrors.NewBizStatusError(5005004, err.Error())
	}
	klog.Infof("Empty cart success for user %d", req.UserId)

	return &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: payResult.TransactionId,
	}, err
}
