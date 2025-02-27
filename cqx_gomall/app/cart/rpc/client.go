package rpc

import (
	"sync"

	"github.com/MrLittle05/Gomall/app/cart/conf"
	cartUtils "github.com/MrLittle05/Gomall/app/cart/utils"
	"github.com/MrLittle05/Gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func Init() {
	once.Do(func() {
		InitProductClient()
	})
}

func InitProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartUtils.MustHandleError(err)
}
