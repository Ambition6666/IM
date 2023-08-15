package sql

import (
	"fmt"
	"im/config"
	"im/internal/service/models"

	mq "github.com/Ambition6666/coderzh.github.io"
	re "github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Rediscc *re.Client
	DB      *gorm.DB
	MQ      *mq.MyRabbitMQ
)

// 注册数据库引擎
func InitSql() {
	InitRedis()
	InitMysql()
	InitMQ()
}

// --------------------------------注册服务----------------------------------
// 注册redis
func InitRedis() {
	addr := config.RedisHost + ":" + config.RedisPort
	fmt.Println(addr)
	Rediscc = re.NewClient(&re.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// 注册mysql
func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.MysqlUser, config.MysqlPwd, config.MysqlHost, config.MysqlPort, config.MysqlDbName)
	fmt.Println(dsn)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}

// 注册MQ
func InitMQ() {
	mq.InitMqConfig(config.RabbitMQUser, config.RabbitMQPwd, config.RabbitMQHost, config.RabbitMQPort)
	MQ = mq.NewRabbitMQSimple("cache")
}

// --------------------------------获取操作服务的句柄----------------------------------
// 得到redis句柄
func GetRedis() *re.Client {
	return Rediscc
}

// 得到mysql的句柄
func GetMysqlDB() *gorm.DB {
	return DB
}

// 得到mq句柄
func GetMQ() *mq.MyRabbitMQ {
	return MQ
}

// --------------------------------初始化表----------------------------------
func RForm() {
	DB.AutoMigrate(&models.Friend_application_list{})
	DB.AutoMigrate(&models.Hail_fellow{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Message{})
}
