package db

import (
	"gorm.io/driver/mysql"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"myblogs/user/grpc/etc"

	"fmt"
)

var MDB *gorm.DB
var RDB *redis.ClusterClient
var err error

func init() {
	var config = etc.NewConfig()
	mysqlDNS := config.Mysql.User + ":" + config.Mysql.Password + "@" + "(" + config.Mysql.Host + ")/" + config.Mysql.Database
	MDB, err = gorm.Open(mysql.Open(mysqlDNS), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to mysql error:", err)
		return
	}

	RDB = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs : config.Redis.Addrs,
	})
}