package main

import (
	"context"
	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	"os"
	"time"

	"github.com/baiyutang/gomall/app/frontend/middleware"
	"github.com/baiyutang/gomall/app/frontend/routes"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"

	"github.com/joho/godotenv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	hertzprom "github.com/hertz-contrib/monitor-prometheus"
	hertzotelprovider "github.com/hertz-contrib/obs-opentelemetry/provider"
	hertzoteltracing "github.com/hertz-contrib/obs-opentelemetry/tracing"

	"github.com/baiyutang/gomall/app/frontend/infra/mtl"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
)

func main() {
	_ = godotenv.Load()

	mtl.InitMtl()
	rpc.InitClient()

	p := hertzotelprovider.NewOpenTelemetryProvider(
		hertzotelprovider.WithSdkTracerProvider(mtl.TracerProvider),
		hertzotelprovider.WithEnableMetrics(false),
	)
	defer p.Shutdown(context.Background())
	tracer, cfg := hertzoteltracing.NewServerTracer()
	h := server.Default(
		server.WithExitWaitTime(time.Second),
		server.WithDisablePrintRoute(false),
		server.WithTracer(
			hertzprom.NewServerTracer(
				"",
				"",
				hertzprom.WithRegistry(mtl.Registry),
				hertzprom.WithDisableServer(true),
			),
		),
		server.WithHostPorts(":8080"),
		tracer,
	)

	store, err := redis.NewStore(100, "tcp", "localhost:6379", "", []byte("AMoIKVVcitM="))
	store.Options(sessions.Options{MaxAge: 86400})
	rs, err := redis.GetRedisStore(store)
	if err == nil {
		rs.SetSerializer(sessions.JSONSerializer{})
	}

	frontendutils.MustHandleError(err)

	h.Use(sessions.New("cloudwego-shop", store))

	h.Use(hertzoteltracing.ServerMiddleware(cfg))
	middleware.RegisterMiddleware(h)

	routes.RegisterProduct(h)
	routes.RegisterHome(h)
	routes.RegisterCategory(h)
	routes.RegisterAuth(h)
	routes.RegisterCart(h)
	routes.RegisterCheckout(h)

	h.LoadHTMLGlob("template/*")
	h.Delims("{{", "}}")
	h.GET("sign-in", func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		session.Set("user_id", 1)
		c.HTML(consts.StatusOK, "sign-in", utils.H{
			"title": "Sign in",
			"next":  c.Query("next"),
		})
	})
	h.GET("sign-up", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "sign-up", utils.H{
			"title": "Sign up",
		})
	})
	h.GET("/order", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "order", utils.H{
			"title": "Order",
		})
	})
	h.GET("/about", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "about", utils.H{
			"title": "About",
		})
	})
	if os.Getenv("GO_ENV") != "online" {
		h.GET("/robots.txt", func(ctx context.Context, c *app.RequestContext) {
			c.Data(consts.StatusOK, "text/plain", []byte(`User-agent: *
Disallow: /`))
		})
	}

	h.Static("/static", "./")
	h.Spin()
}
