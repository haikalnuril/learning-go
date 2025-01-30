package config

import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectDB() (*gorm.DB, error) {
	dsn := databaseCredential()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

func databaseCredential() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	if port == "" {
		port = "5432"
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
}