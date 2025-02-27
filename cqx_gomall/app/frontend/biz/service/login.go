package service

import (
	"context"

	"github.com/MrLittle05/Gomall/app/frontend/infra/rpc"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/user"
	"github.com/google/uuid"

	auth "github.com/MrLittle05/Gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (data map[string]any, err error) {
	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		hlog.Errorf("Login service failed for user %s: %v", req.Email, err)
		return nil, err
	}

	// 将sessionId关联用户信息存入redis
	sessionId := uuid.New().String()
	session := sessions.Default(h.RequestContext)
	session.Set(sessionId, resp.UserId)
	err = session.Save()
	if err != nil {
		hlog.Errorf("Save session failed for user %d: %v", resp.UserId, err)
		return nil, err
	}

	redirect := "/"
	if req.Next != "" {
		redirect = req.Next
	}
	return utils.H{
		"redirect":  redirect,
		"sessionId": sessionId,
	}, nil
}
