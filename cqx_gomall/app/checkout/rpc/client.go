package rpc

import (
	"sync"

	"github.com/MrLittle05/Gomall/app/checkout/conf"
	"github.com/MrLittle05/Gomall/app/checkout/utils"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	OrderClient   orderservice.Client
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PayClient     paymentservice.Client
	once          sync.Once
)

func Init() {
	once.Do(func() {
		InitCartClient()
		InitPayClient()
		InitProductClient()
		InitOrderClient()
	})
}

func InitCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	utils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart", opts...)
	utils.MustHandleError(err)
}

func InitPayClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	utils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	PayClient, err = paymentservice.NewClient("payment", opts...)
	utils.MustHandleError(err)
}

func InitProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	utils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	utils.MustHandleError(err)
}

func InitOrderClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	utils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	OrderClient, err = orderservice.NewClient("order", opts...)
	utils.MustHandleError(err)
}
