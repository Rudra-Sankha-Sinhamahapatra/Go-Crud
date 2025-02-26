package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort int
	DBURL      string
	JWT_SECRET string
}

var AppConfig Config

func LoadEnv() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
		return Config{}, err
	}

	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)

	if err != nil || port <= 0 {
		port = 8900
	}

	db := os.Getenv("DATABASE_URL")

	if db == "" {
		db = "postgresql://postgres:mysecretpassword@localhost:5432/postgres"
	}

	JWT_SECRET := os.Getenv("JWT_SECRET")

	if JWT_SECRET == "" {
		JWT_SECRET = "SECRET"
	}

	AppConfig = Config{
		ServerPort: port,
		DBURL:      db,
		JWT_SECRET: JWT_SECRET,
	}

	return AppConfig, nil
}
