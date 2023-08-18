package db

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ThreadId    string `gorm:"column:thread_id;type:varchar(255);not null"`
	FromUserId  int64  `gorm:"column:from_user_id;not null;index:fk_user_message_from"`
	ToUserId    int64  `gorm:"column:to_user_id;not null;index:fk_user_message_to"`
	Contents    string `gorm:"column:contents;type:varchar(255);not null"`
	MessageUUID int64  `gorm:"column:message_uuid;not null;index:fk_uuid_message"`
	CreateTime  int64  `gorm:"column:create_time;not null;"`
}

func (Message) TableName() string {
	return "message"
}
