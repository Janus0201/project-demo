package dal

import (
	"github.com/MrLittle05/Gomall/app/cart/biz/dal/mysql"
	// "github.com/MrLittle05/Gomall/app/cart/biz/dal/redis"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
