package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDB() {
	connectionStr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", ENV.DB_HOST, ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_NAME, ENV.DB_PORT)
	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}
