package service

import (
	"context"

	auth "github.com/MrLittle05/Gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/MrLittle05/Gomall/app/frontend/infra/rpc"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *auth.Empty, err error) {
	_, err = rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	})
	if err != nil {
		return nil, err
	}

	return
}
