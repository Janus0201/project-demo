package auth

import (
	"context"

	"github.com/MrLittle05/Gomall/app/frontend/biz/service"
	"github.com/MrLittle05/Gomall/app/frontend/biz/utils"
	auth "github.com/MrLittle05/Gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
)

// Login .
// @router /auth/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	data, err := service.NewLoginService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// Set cookie
	c.SetCookie("session_id", data["sessionId"].(string), 3600, "/", "localhost", protocol.CookieSameSiteLaxMode, false, true)

	c.Redirect(consts.StatusOK, []byte(data["redirect"].(string)))
}

// Register .
// @router /auth/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewRegisterService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.Redirect(consts.StatusOK, []byte("/sign-in"))
}

// Logout .
// @router /auth/logout [POST]
func Logout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewLogoutService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// Clear session
	sessionID := string(c.Cookie("session_id"))
	session := sessions.Default(c)
	session.Delete(sessionID)
	err = session.Save()
	if err != nil {
		hlog.Errorf("Logout(): Failed to save session: %v", err)
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// Clear cookie
	c.SetCookie("session_id", "", -1, "/", "localhost", protocol.CookieSameSiteLaxMode, false, false)

	c.Redirect(consts.StatusOK, []byte("/"))
}
