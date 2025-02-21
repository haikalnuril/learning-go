package dto

import "mime/multipart"

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PostRequest struct {
	UserID  uint                  `form:"user_id"`
	Tweet   string                `form:"tweet"`
	Picture *multipart.FileHeader `form:"picture"`
}


type PostResponse struct {
	ID         uint   `json:"id"`
	UserID     uint   `json:"-"`
	User       User   `gorm:"foreignKey:UserID"`
	Tweet      string `json:"tweet"`
	PictureUrl string `json:"picture_url"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}