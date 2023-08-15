package config

import (
	"os"

	"github.com/spf13/viper"
)

// 声明配置
var (
	ServerPort   string
	ServerHost   string
	RedisPort    string
	RedisHost    string
	MysqlPort    string
	MysqlHost    string
	MysqlPwd     string
	MysqlDbName  string
	MysqlUser    string
	ConsulHost   string
	ConsulPort   string
	RabbitMQHost string
	RabbitMQPort string
	RabbitMQUser string
	RabbitMQPwd  string
)

// 注册配置
func ConfigInit() {
	InitConfig()
	ServerPort = viper.GetString("Server.port")
	ServerHost = viper.GetString("Server.host")
	RedisHost = viper.GetString("Redis.host")
	RedisPort = viper.GetString("Redis.Port")
	MysqlHost = viper.GetString("Mysql.host")
	MysqlPort = viper.GetString("Mysql.port")
	MysqlPwd = viper.GetString("Mysql.password")
	MysqlDbName = viper.GetString("Mysql.dbname")
	MysqlUser = viper.GetString("Mysql.user")
	ConsulHost = viper.GetString("Consul.host")
	ConsulPort = viper.GetString("Consul.port")
	RabbitMQHost = viper.GetString("RabbitMQ.host")
	RabbitMQPort = viper.GetString("RabbitMQ.port")
	RabbitMQUser = viper.GetString("RabbitMQ.user")
	RabbitMQPwd = viper.GetString("RabbitMQ.password")
}

// 获取配置
func InitConfig() {
	workdir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workdir + "/config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
