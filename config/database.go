package config

import (
	"log"

	"fixfon/go-telegram-bot-boilerplate/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(postgres.Open(AppConfig.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("[ConnectDB()] Failed to connect to database:", err)
	}

	// Run migrations
	if err := migrations.RunMigrations(DB); err != nil {
		log.Fatal("[ConnectDB()] Failed to run migrations:", err)
	}

	log.Println("[ConnectDB()] Database migrations completed successfully!")
	DB = db
}
