package services

import (
	"taskman/internal/models"
	"taskman/internal/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserById(id uint) (*models.User, error)
	UpdateUser(id uint, user *models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserById(id uint) (*models.User, error) {
	return s.repo.GetUserById(id)
}

func (s *userService) UpdateUser(id uint, user *models.User) error {
	return s.repo.UpdateUser(id, user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}