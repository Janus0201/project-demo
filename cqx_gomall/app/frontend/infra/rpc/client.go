package rpc

import (
	"sync"

	"github.com/MrLittle05/Gomall/app/frontend/conf"
	frontendUtils "github.com/MrLittle05/Gomall/app/frontend/utils"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	OrderClient    orderservice.Client
	CheckoutClient checkoutservice.Client
	CartClient     cartservice.Client
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	once           sync.Once
)

func Init() {
	once.Do(func() {
		InitUserClient()
		InitProductClient()
		InitCartClient()
		InitCheckoutClient()
		InitOrderClient()
	})
}

// 注册用户服务UserService的客户端UserClient，建立请求一个RPC连接，实现远程调用
func InitUserClient() {
	// 服务发现
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	// 如果没有 ...，你需要手动将切片中的每个元素逐一传递，这会非常繁琐。使用 ... 可以自动展开切片，简化代码
	UserClient, err = userservice.NewClient("user", opts...)
	frontendUtils.MustHandleError(err)
}

func InitProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendUtils.MustHandleError(err)
}

func InitCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendUtils.MustHandleError(err)
}

func InitCheckoutClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	CheckoutClient, err = checkoutservice.NewClient("checkout", opts...)
	frontendUtils.MustHandleError(err)
}

func InitOrderClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	OrderClient, err = orderservice.NewClient("order", opts...)
	frontendUtils.MustHandleError(err)
}
