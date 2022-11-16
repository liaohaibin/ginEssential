package model

import (
	"github.com/jinzhu/gorm"
)

type Pub struct {
	gorm.Model
	User string `gorm:"type:varchar(20);not null"`
	Project string `gorm:"type:varchar(110);not null"`
	DestHost  string `gorm:"size:255;not null"`
}