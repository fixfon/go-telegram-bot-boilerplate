package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
		SSLMode  string
	}
	TelegramToken string
}

var AppConfig Config

func LoadConfig() error {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println("[LoadConfig()] error loading .env file will use environment variables: " + err.Error())
	}

	// Load configuration from environment variables
	AppConfig.Database.Host = getEnv("APP_DATABASE_HOST", "localhost")
	AppConfig.Database.Port = getEnv("APP_DATABASE_PORT", "5432")
	AppConfig.Database.User = getEnv("APP_DATABASE_USER", "postgres")
	AppConfig.Database.Password = getEnv("APP_DATABASE_PASSWORD", "")
	AppConfig.Database.DBName = getEnv("APP_DATABASE_NAME", "telegram_bot")
	AppConfig.Database.SSLMode = getEnv("APP_DATABASE_SSLMODE", "disable")

	// Telegram token is required
	AppConfig.TelegramToken = os.Getenv("APP_TELEGRAM_TOKEN")
	if AppConfig.TelegramToken == "" {
		return errors.New("[LoadConfig()] telegram token is required")
	}

	return nil
}

// getEnv gets environment variable with a default fallback value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
