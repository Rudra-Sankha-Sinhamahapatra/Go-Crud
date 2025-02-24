package utils

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	if AppConfig.DBURL == "" {
		log.Fatal("database URL is not set. Make sure LoadEnv() runs first.")
	}

	DB, err = gorm.Open(postgres.Open(AppConfig.DBURL), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Database connected Successfully")

}
