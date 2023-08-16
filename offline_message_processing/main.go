package main

//通过mq后台处理离线消息
import (
	"encoding/json"
	"fmt"
	"im/internal/service/models"
	"im/vo"

	mq "github.com/Ambition6666/coderzh.github.io"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// config.ConfigInit()
	InitMysql()
}
func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "root", "192.168.40.132", "3306", "chat")
	fmt.Println(dsn)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	mq.InitMqConfig("guest", "guest", "192.168.40.132", "5672")
	m := mq.NewRabbitMQSimple("cache")
	mes := m.ConsumeSimple()
	//db := sql.GetMysqlDB()
	go func() {
		for v := range mes {
			mes1 := new(vo.Message)
			err := json.Unmarshal(v.Body, mes1)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(mes1)
			ms := &models.Message{
				FromId:     mes1.Fr,
				TargetId:   mes1.To,
				Context:    mes1.Ctx,
				Type:       mes1.T,
				TargetType: mes1.TT,
			}
			DB.Create(ms)
			// repository.InsertMessage(ms)
			// fmt.Printf("%s\n", v.Body)
		}
	}()
	select {}
}
