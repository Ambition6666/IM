package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromId     string
	TargetId   string
	Context    string
	Type       uint
	TargetType uint
}

/*
type:
1-->text(普通文本)
2-->picture(图片)
3-->view(视频)
4-->music(音乐)
-------------------------------------------------
target_type:
1-->private(私人聊天)
2-->group(群聊)
3-->boardcast(广播)
*/
