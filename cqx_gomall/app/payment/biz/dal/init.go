package dal

import (
	"github.com/MrLittle05/Gomall/app/payment/biz/dal/mysql"
	// "github.com/MrLittle05/Gomall/app/payment/biz/dal/redis"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
