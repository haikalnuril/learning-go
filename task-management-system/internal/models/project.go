package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name 		string `gorm:"not null"`
	Description string `gorm:"not null"`
	Status 		string `gorm:"type:varchar(20);not null;check:status IN ('pending', 'in-progress', 'done')"`
	UserID 		uint
	User 		User `gorm:"foreignKey:UserID"`
}