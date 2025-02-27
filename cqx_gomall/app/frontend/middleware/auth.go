package middleware

import (
	"context"
	"net/url"
	"path"
	"strings"

	"github.com/MrLittle05/Gomall/app/frontend/conf"
	frontendUtils "github.com/MrLittle05/Gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sessions"
)

/*
日志级别区分：
ERROR通常表示发生了错误，导致某个操作无法完成，而WARN是问题存在但还未影响主要功能。
例如，配置文件缺失但使用了默认值，这时用WARN；而如果无法连接到关键数据库导致功能不可用，则是ERROR。
*/

// 每个HTTP请求都对应一个上下文 ctx, 把用户ID保存到上下文中

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		key := string(c.Cookie("session_id"))
		if key == "" {
			hlog.Error("GlobalAuth(): Cookie key is empty")
			c.Next(ctx)
			return
		}
		session := sessions.Default(c)
		userId := session.Get(key)
		//		hlog.Info("userid from session: ", userId)
		if userId == nil {
			hlog.Error("GlobalAuth(): userid is nil")
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, frontendUtils.SessionUserIdKey, userId)
		//		hlog.Info("userId in context: ", ctx.Value(frontendUtils.SessionUserIdKey))
		c.Next(ctx)
	}
}

// 用户鉴权中间件
func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		requestPath := c.FullPath()
		if isWhitelisted(requestPath) {
			hlog.Info("Request path is whitelisted, skipping Auth middleware")
			c.Next(ctx)
			return
		}

		key := string(c.Cookie("session_id"))
		// cookie不存在，说明用户尚未登录，跳转到登录页面
		if key == "" {
			hlog.Info("No session_id cookie, redirecting to login")
			redirectToLogin(c, requestPath)
			return
		}

		session := sessions.Default(c)
		userId := session.Get(key)
		//		hlog.Info("userid from session: ", userId)
		// 用户不存在，说明用户尚未登录，跳转到登录页面
		if userId == nil {
			hlog.Info("Session exists but userID is missing, possible session tampering")
			redirectToLogin(c, requestPath)
			return
		}

		ctx = context.WithValue(ctx, frontendUtils.SessionUserIdKey, userId)
		c.Next(ctx)
	}
}

// 第二层编码：url.QueryEscape 对合法路径中的其他特殊字符（如 /, =）进行编码，确保传输安全。
func redirectToLogin(c *app.RequestContext, next string) {
	if !isValidPath(next) {
		hlog.Warn("Invalid redirect path detected")
		next = "/"
	}
	encodedNext := url.QueryEscape(next)
	c.Redirect(302, []byte("/sign-in?next="+encodedNext))
	c.Abort()
}

// 第一层检查：isValidPath 函数拒绝包含 ? 和 # 的路径（从源头拦截非法格式）。
func isValidPath(path string) bool {
	return strings.HasPrefix(path, "/") && !strings.ContainsAny(path, "#?")
}

func isWhitelisted(requestPath string) bool {
	whitelist := conf.GetConf().AuthWhitelist
	for _, p := range whitelist {
		// 使用 path.Match 支持通配符
		match, _ := path.Match(p, requestPath)
		if match {
			return true
		}
	}
	return false
}
