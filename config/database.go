package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	var err error

	DB, err = gorm.Open(postgres.Open(AppConfig.DatabaseURL), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("[ConnectDB()] Database connected successfully!")
	return nil
}

// Get database instance
func GetDB() *gorm.DB {
	return DB
}
