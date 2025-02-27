package service

import (
	"context"

	checkout "github.com/MrLittle05/Gomall/app/frontend/hertz_gen/frontend/checkout"
	"github.com/MrLittle05/Gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/MrLittle05/Gomall/app/frontend/utils"
	rpccheckout "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/checkout"
	payment "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    uint32(userId),
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Phone:     req.Email,
		Address: &rpccheckout.Address{
			StreetAddress: req.Street,
			Country:       req.Country,
			State:         req.Province,
			City:          req.City,
			ZipCode:       req.Zipcode,
		},
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CardNum,
			CreditCardExpirationMonth: req.ExpirationMonth,
			CreditCardExpirationYear:  req.ExpirationYear,
			CreditCardCvv:             req.Cvv,
		},
	})
	if err != nil {
		hlog.Errorf("Checkout service failed for user %d: %v", userId, err)
		return utils.H{
			"error": true,
		}, nil
	}

	hlog.Infof("Checkout service success for user %d", userId)
	return utils.H{
		"redirect": true,
	}, nil
}
