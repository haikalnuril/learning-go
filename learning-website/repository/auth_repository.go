package repository

import (
	"gin-socmed/model"

	"gorm.io/gorm"
)

type AuthRepository interface{
	EmailExist(email string) bool
	Register(user *model.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) EmailExist(email string) bool {
	var user model.User
	err := r.db.First(&user, "Email=?", email).Error

	return err == nil
}

func (r *authRepository) Register(user *model.User) error {
	err := r.db.Create(&user).Error

	return err
}