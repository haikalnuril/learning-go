package main

import (
	"log"
	"taskman/internal/config"
	"taskman/internal/models"
	"github.com/joho/godotenv"
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

	
	err = db.AutoMigrate(&models.User{}, &models.Project{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}
	log.Println("Database migrated successfully")
}