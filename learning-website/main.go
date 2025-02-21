package main

import (
	"fmt"
	"gin-socmed/config"
	"gin-socmed/model"
	"gin-socmed/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.LoadDB()

	config.DB.AutoMigrate(&model.User{}, &model.Post{})

	r := gin.Default()
	api := r.Group("/api/v1")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.AuthRouter(api)
	router.PostRouter(api)

	r.Run(fmt.Sprintf(":%v", config.ENV.PORT))
}