package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Gender   string	`gorm:"type:varchar(10);not null;check:gender IN ('male', 'female');default:'male'"`
}