package model

import "github.com/jinzhu/gorm"

type relation struct {
	gorm.Model
	UserId   int `gorm:"type:int;not null" json:"user_id"`
	FollowId int `gorm:"type:int;not null" json:"follow_id"`
}
