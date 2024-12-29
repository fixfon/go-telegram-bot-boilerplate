package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL   string
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
	AppConfig.DatabaseURL = getEnv("APP_DATABASE_URL", "")
	if AppConfig.DatabaseURL == "" {
		return errors.New("[LoadConfig()] database URL is required")
	}

	// Telegram token is required
	AppConfig.TelegramToken = getEnv("APP_TELEGRAM_TOKEN", "")
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
