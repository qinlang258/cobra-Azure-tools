package database

//https://gorm.io/zh_CN/docs/connecting_to_the_database.html

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ContentMysql() *gorm.DB {
	//读取.ini里面的数据库配置

	dsn := fmt.Sprintf("root:ql2252528@tcp(192.168.44.134:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields:            true, //打印sql
		SkipDefaultTransaction: true,
	})
	// DB.Debug()
	if err != nil {
		fmt.Println(err)
	}
	return DB
}
