// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	auth "github.com/MrLittle05/Gomall/app/frontend/biz/router/auth"
	cart "github.com/MrLittle05/Gomall/app/frontend/biz/router/cart"
	category "github.com/MrLittle05/Gomall/app/frontend/biz/router/category"
	checkout "github.com/MrLittle05/Gomall/app/frontend/biz/router/checkout"
	home "github.com/MrLittle05/Gomall/app/frontend/biz/router/home"
	order "github.com/MrLittle05/Gomall/app/frontend/biz/router/order"
	product "github.com/MrLittle05/Gomall/app/frontend/biz/router/product"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	order.Register(r)

	checkout.Register(r)

	cart.Register(r)

	category.Register(r)

	product.Register(r)

	auth.Register(r)

	home.Register(r)
}
