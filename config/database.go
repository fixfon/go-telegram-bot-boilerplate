package config

import (
	"fmt"
	"log"

	"fixfon/go-telegram-bot-boilerplate/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		AppConfig.Database.Host,
		AppConfig.Database.Port,
		AppConfig.Database.User,
		AppConfig.Database.Password,
		AppConfig.Database.DBName,
		AppConfig.Database.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	if err := migrations.RunMigrations(DB); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	fmt.Println("Database migrations completed successfully!")
	DB = db
}
