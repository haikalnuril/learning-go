package handler

import (
	"gin-socmed/dto"
	"gin-socmed/errorhandler"
	"gin-socmed/helper"
	"gin-socmed/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var register dto.RegisterRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		errorhandler.HandlerError(c, &errorhandler.BadRequestError{Message: err.Error()})

		return
	}

	dataRegister, err := h.service.Register(&register)
	if err != nil {
		errorhandler.HandlerError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Register Successfully",
		Data:       dataRegister,
	})

	c.JSON(http.StatusCreated, res)
}


func (h *authHandler) Login(c *gin.Context) {
	var login dto.LoginRequest

	err := c.ShouldBindJSON(&login)

	if err != nil {
		errorhandler.HandlerError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	result, err := h.service.Login(&login)
	if err != nil {
		errorhandler.HandlerError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message: "Success to login",
		Data: result,
	})

	c.JSON(http.StatusOK, res)
}