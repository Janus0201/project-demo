package middleware

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(h *server.Hertz) {
	// Use 方法用于将全局中间件附加到路由器上。这意味着该中间件会对所有通过此路由器处理的 HTTP 请求生效
	h.Use(GlobalAuth())
	h.Use(Auth())
}
