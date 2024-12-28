package config

import (
	"fmt"
	"log"

	"fixfon/wallet_telegram/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=your_password dbname=telegram_bot port=5432 sslmode=disable"
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
