package main

import (
	"crud/src/models"
	"crud/src/utils"
	"log"
)

func Migrate() {
	if _, err := utils.LoadEnv(); err != nil {
		log.Fatal("Failed to load environment:", err)
	}
	utils.InitDB()

	if err := utils.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration completed successfully")
}

func main() {
	Migrate()
}
