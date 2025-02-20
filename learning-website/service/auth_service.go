package service

import (
	"gin-socmed/dto"
	"gin-socmed/errorhandler"
	"gin-socmed/helper"
	"gin-socmed/model"
	"gin-socmed/repository"
)

type AuthService interface{
	Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login (req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	if emailExist := s.repository.EmailExist(req.Email); emailExist {
		return nil, &errorhandler.BadRequestError{Message: "email already registered"}
	}

	if req.Password != req.PasswordConfirmation {
		return nil, &errorhandler.BadRequestError{Message: "password not match"}
	}

	passwordHash, err := helper.HashPassword(req.Password);

	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	user := model.User {
		Name: req.Name,
		Email: req.Email,
		Password: passwordHash,
		Gender: req.Gender,
	}

	if err := s.repository.Register(&user); err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	return &dto.RegisterResponse{
		Name: user.Name,
		Email: user.Email,
		Password: passwordHash,
		Gender: user.Gender,
	}, nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse

	user, err := s.repository.GetUserByEmail(req.Email)

	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "wrong email or password"}
	}

	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorhandler.NotFoundError{Message: "wrong email or password"}
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	data = dto.LoginResponse{
		ID: int(user.ID),
		Name: user.Name,
		Email: user.Email,
		Token: token,
	}

	return &data, nil
}