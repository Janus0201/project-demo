package dal

import (
	"github.com/MrLittle05/Gomall/app/user/biz/dal/mysql"
	"github.com/MrLittle05/Gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
