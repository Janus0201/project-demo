package mysql

import (
	"fmt"
	"os"

	"github.com/MrLittle05/Gomall/app/order/biz/model"
	"github.com/MrLittle05/Gomall/app/order/conf"
	"github.com/cloudwego/kitex/pkg/klog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		klog.Errorf("Failed to connect to MySQL DB: %v", err)
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		err = DB.AutoMigrate(&model.Order{}, &model.OrderItem{})
		if err != nil {
			klog.Errorf("Failed to migrate MySQL DB: %v", err)
			panic(err)
		}
	}
}
