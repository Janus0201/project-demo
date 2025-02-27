package dal

import (
	"github.com/MrLittle05/Gomall/app/frontend/biz/dal/mysql"
	"github.com/MrLittle05/Gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
