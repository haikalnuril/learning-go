package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Tweet      string  `gorm:"not null;type:text"`
	PictureUrl *string `gorm:"type:text"`
	UserID     uint
	User       User `gorm:"foreignKey:UserID"`
}
