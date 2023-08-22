package etc

import (
	"github.com/spf13/viper"
	"fmt"
)

type Config struct {
	ListenOn grpcConfig
	Mysql mysqlConfig
}

type grpcConfig struct {
	Host string
	Port int
}

type mysqlConfig struct {
	Host string
	User string
	Password string
	Database string
	Port int
}

var (
	My_Config *Config = NewConfig()
)

func NewConfig() *Config {
	var config Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
    viper.AddConfigPath("./etc")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic("loading config file error!")
	}

	viper.Unmarshal(&config)
	return &config
}
