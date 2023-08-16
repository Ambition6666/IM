package repository

import (
	"im/internal/service/models"
	"im/sql"
)

//-------------------------- 对消息进行操作-------------------------------

// 插入消息
func InsertMessage(a *models.Message) {
	db := sql.GetMysqlDB()
	db.Create(a)
}

// 若发送消息对方不在线则存入mq
func InsertCache(a string) {
	mq := sql.GetMQ()
	mq.PublishSimple(a)
}

// 获取聊天记录,第一个是本人,第二个是他人
func GetMessages(n string, s string) []models.Message {
	db := sql.GetMysqlDB()
	v := make([]models.Message, 0)
	db.Where("from_id = ? and target_id = ?", s, n).Find(&v)
	return v
}

// 获取群聊记录
func GetGroupMessages(g string) []models.Message {
	db := sql.GetMysqlDB()
	v := make([]models.Message, 0)
	db.Where("target_id = ?", g).Find(&v)
	return v
}
