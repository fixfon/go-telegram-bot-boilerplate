package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL   string
	TelegramToken string
	OwnerID       int64
}

var AppConfig Config

func LoadConfig() error {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println("[LoadConfig()] error loading .env file will use environment variables: " + err.Error())
	}

	AppConfig.DatabaseURL = os.Getenv("APP_DATABASE_URL")
	if AppConfig.DatabaseURL == "" {
		return errors.New("[LoadConfig()] database URL is required")
	}

	// Telegram token is required
	AppConfig.TelegramToken = os.Getenv("APP_TELEGRAM_TOKEN")
	if AppConfig.TelegramToken == "" {
		return errors.New("[LoadConfig()] telegram token is required")
	}

	// Parse owner ID from env
	ownerID := os.Getenv("APP_OWNER_ID")
	if ownerID == "" {
		return errors.New("[LoadConfig()] owner ID is required")
	}

	// Convert string to int64
	id, err := strconv.ParseInt(ownerID, 10, 64)
	if err != nil {
		return fmt.Errorf("[LoadConfig()] invalid owner ID: %v", err)
	}
	AppConfig.OwnerID = id

	return nil
}
