package mysql

import (
	"fmt"
	"github.com/MrLittle05/Gomall/app/user/biz/model"
	"github.com/MrLittle05/Gomall/app/user/conf"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	// DSN: Data Source Name，具体来说，这是一个用于数据库连接的字符串，特别针对使用 GORM 与 MySQL 数据库进行交互的情况
	/* 使用 fmt.Sprintf 函数结合环境变量来构建 DSN（数据源名称）字符串是一种常见的做法，目的是提高配置的灵活性和安全性。
	通过这种方式，可以避免将敏感信息如数据库用户名、密码等直接硬编码在代码或配置文件中，而是从系统的环境变量中读取这些值。*/
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	/* 自动迁移可以帮助开发者快速地保持数据库模式与 Go 代码中定义的模型同步。
	当你添加新的字段、修改字段类型或删除字段时，自动迁移能够帮助你更新数据库结构，使得数据库结构与Go模型一致。*/
	DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
