package dal

import (
	"github.com/MrLittle05/Gomall/app/product/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
