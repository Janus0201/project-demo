package dal

import (
	"github.com/MrLittle05/Gomall/app/order/biz/dal/mysql"
	// "github.com/MrLittle05/Gomall/app/order/biz/dal/redis"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
