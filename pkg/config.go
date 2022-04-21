package pkg

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type AppConfig struct {
	AppName string
	DB      DatabaseConfig
}

func InitConfig() *AppConfig {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Print("Error loading .env file")
	}

	return &AppConfig{
		AppName: os.Getenv("APP_NAME"),
		DB: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_DATABASE"),
		},
	}
}
