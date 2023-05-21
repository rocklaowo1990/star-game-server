package model

import "gorm.io/gorm"

type InstantMessage struct {
	gorm.Model

	FromUid   string `json:"from"`    // 发送着
	TargetUid string `json:"target"`  // 接收着
	Type      int    `json:"type"`    // 洗哦嘻类型 群聊私聊
	Media     int    `json:"media"`   // 消息类型 文字图片饮品等
	Content   string `json:"content"` // 消息内容
}

func (InstantMessage) TableName() string {
	return "instant_message"
}
