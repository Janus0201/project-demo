package service

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/MrLittle05/Gomall/app/checkout/rpc"
	checkout "github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/checkout"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/payment"
)

func TestCheckout_Run(t *testing.T) {
	fmt.Println(os.Getwd())
	rpc.Init()
	ctx := context.Background()
	s := NewCheckoutService(ctx)
	// init req and assert value

	req := &checkout.CheckoutReq{
		UserId:    1,
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "1234567890",
		Address: &checkout.Address{
			StreetAddress: "123 Main St",
			Country:       "USA",
			State:         "CA",
			City:          "Anytown",
			ZipCode:       "12345",
		},
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          "1234567890123456",
			CreditCardExpirationMonth: 12,
			CreditCardExpirationYear:  2026,
			CreditCardCvv:             123,
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
