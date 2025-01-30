package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name 		string `gorm:"not null"`
	Email 		string `gorm:"not null;unique"`
	Password 	string `gorm:"not null"`
	photo	 	string `gorm:"default:null"`
	Projects 	[]Project
}