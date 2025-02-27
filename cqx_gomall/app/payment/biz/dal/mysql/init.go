package mysql

import (
	"fmt"
	"os"

	"github.com/MrLittle05/Gomall/app/payment/biz/model"
	"github.com/MrLittle05/Gomall/app/payment/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if conf.GetConf().Env != "online" {
		DB.AutoMigrate(&model.PaymentLog{})
	}
}
