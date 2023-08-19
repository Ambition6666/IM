package models

import (
	"gorm.io/gorm"
)

// 存储用户
type User struct {
	gorm.Model
	Name  string `gorm:"unique;not null;type:varchar(200)"`
	Pwd   string
	Email string
	Phone string
}

// 好友
type Hail_fellow struct {
	gorm.Model
	Name         string
	FellowName   string
	Last_message string
}

// 好友申请
type Friend_application_list struct {
	Name string
	gorm.Model
	FromId string
}
