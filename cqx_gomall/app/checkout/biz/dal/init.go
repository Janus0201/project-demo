package dal

import (
	"github.com/MrLittle05/Gomall/app/checkout/biz/dal/mysql"
	"github.com/MrLittle05/Gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
