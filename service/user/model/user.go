package model

import (
	"github.com/jinzhu/gorm"
)

type user struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;unique" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
}
