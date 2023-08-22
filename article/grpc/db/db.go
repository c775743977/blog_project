package db

import (
	"gorm.io/driver/mysql"

	"gorm.io/gorm"

	"grpc_article/etc"

	"fmt"
)

var MDB *gorm.DB
var err error

func init() {
	// var config = etc.NewConfig()
	mysqlDNS := etc.My_Config.Mysql.User + ":" + etc.My_Config.Mysql.Password + "@" + "(" + etc.My_Config.Mysql.Host + ")/" + etc.My_Config.Mysql.Database
	MDB, err = gorm.Open(mysql.Open(mysqlDNS), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to mysql error:", err)
		return
	}
}