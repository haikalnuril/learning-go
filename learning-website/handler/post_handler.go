package handler

import (
	"fmt"
	"gin-socmed/dto"
	"gin-socmed/errorhandler"
	"gin-socmed/helper"
	"gin-socmed/service"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type postHandler struct {
	service service.PostService
}

func NewPostHandler(s service.PostService) *postHandler {
	return &postHandler{
		service: s,
	}
}

func (h *postHandler) Create(c *gin.Context) {
	var post dto.PostRequest

	if err := c.ShouldBind(&post); err != nil {
		errorhandler.HandlerError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if post.Picture != nil {
		if err := os.MkdirAll("/public/picture", 0755); err != nil {
			errorhandler.HandlerError(c, &errorhandler.InternalServerError{Message: err.Error()})
			return
		}

		//rename picture

		ext := filepath.Ext(post.Picture.Filename)
		newFileName := uuid.New().String() + ext

		//save image to directory

		dst := filepath.Join("/public/picture", filepath.Base(newFileName))
		c.SaveUploadedFile(post.Picture, dst)

		post.Picture.Filename = fmt.Sprintf("%s/public/picture/%s", c.Request.Host, newFileName)
	}

	userID, _ := c.Get("userID")
	post.UserID = userID.(uint)

	if err := h.service.Create(&post); err != nil {
		errorhandler.HandlerError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message: "Success to Create Tweet",
	})

	c.JSON(http.StatusCreated, res)
}
