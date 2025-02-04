package handlers

import (
	"net/http"
	"taskman/internal/models"
	"taskman/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler (service services.UserService) *UserHandler {
	return &UserHandler{service:service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H(
			"status": "Failed",
			"error": "invalid request. Name, Email, and Password are required",
			))
	}

	if err := h.service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError), gin.H(
			"status": "Failed",
			"error": "Internal Server Error. Can't Create User"
		)
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error": "Internal Server Error. Can't Get Users"
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "Users Found Successfully",
		"data": users
	})
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Failed",
			"error": "User Not Found"
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "User Found Successfully",
		"data": user
		}
	)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed",
			"error": "Invalid Request. Name, Email, and Password are required"
		})
		return
	}

	user.ID = id

	if err := h.service.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error": "Internal Server Error. Can't Update User"
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "User Updated Successfully",
		"data": user
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error": "Internal Server Error. Can't Delete User"
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "User Deleted Successfully"
	})
}