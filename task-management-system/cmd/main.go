package main

import (
	"log"
	"taskman/internal/config"
	"taskman/internal/handlers"
	"taskman/internal/services"
	"taskman/internal/repositories"
	"taskman/internal/models"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	log.Println("Database connected successfully")

	
	err = db.AutoMigrate(&models.User{}, &models.Project{}, &models.Task{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}
	log.Println("Database migrated successfully")

	userRepo := repositories.NewUserRepository(db)
	projectRepo := repositories.NewProjectRepository(db)
	taskRepo := repositories.NewTaskRepository(db)

	userService := services.NewUserService(userRepo)
	projectService := services.NewProjectService(projectRepo)
	taskService := services.NewTaskService(taskRepo)

	userHandler := handlers.NewUserHandler(userService)
	projectHandler := handlers.NewProjectHandler(projectService)
	taskHandler := handlers.NewTaskHandler(taskService)

	r := gin.Default()

	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetAllUsers)
	r.GET("/users/:id", userHandler.GetUserById)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	r.POST("/projects", projectHandler.CreateProject)
	r.GET("/projects", projectHandler.GetAllProjects)
	r.GET("/projects/:id", projectHandler.GetProjectById)
	r.PUT("/projects/:id", projectHandler.UpdateProject)
	r.DELETE("/projects/:id", projectHandler.DeleteProject)

	r.POST("/tasks", taskHandler.CreateTask)
	r.GET("/tasks", taskHandler.GetAllTasks)
	r.GET("/tasks/:id", taskHandler.GetTaskById)
	r.PUT("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)

	r.Run(":3000")
}

