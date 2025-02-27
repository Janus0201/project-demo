package service

import (
	"context"
	"strconv"
	"time"

	"github.com/MrLittle05/Gomall/app/payment/biz/dal/mysql"
	"github.com/MrLittle05/Gomall/app/payment/biz/model"
	payment "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}

	// Validate the card
	err = card.Validate(true)
	if err != nil {
		klog.Errorf("Credit card validation failed for user %d, order %s: %v", req.UserId, req.OrderId, err)
		return nil, kerrors.NewGRPCBizStatusError(4004001, err.Error())
	}

	// UUID 的设计确保了全球范围内的唯一性，即使是分布式系统中多个节点同时生成 UUID，冲突的概率也非常低
	transactionId, err := uuid.NewRandom()
	if err != nil {
		klog.Errorf("Failed to generate UUID for user %d, order %s: %v", req.UserId, req.OrderId, err)
		return nil, kerrors.NewGRPCBizStatusError(4005001, err.Error())
	}

	// Create payment log
	err = model.CreatePaymentLog(s.ctx, mysql.DB, &model.PaymentLog{
		UserID:        req.UserId,
		OrderID:       req.OrderId,
		TransactionId: transactionId.String(),
		Amount:        float64(req.Amount),
		PayAt:         time.Now(),
	})
	if err != nil {
		klog.Errorf("Failed to create payment log for user %d, order %s: %v", req.UserId, req.OrderId, err)
		return nil, kerrors.NewGRPCBizStatusError(4005002, err.Error())
	}

	klog.Infof("Payment log created for user %d, order %s", req.UserId, req.OrderId)
	return &payment.ChargeResp{TransactionId: transactionId.String()}, err
}
