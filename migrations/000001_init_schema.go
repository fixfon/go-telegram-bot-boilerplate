package migrations

import (
	"fixfon/go-telegram-bot-boilerplate/models"
	"gorm.io/gorm"
)

// Migration represents a single database migration
type Migration struct {
	ID      uint   `gorm:"primaryKey"`
	Version string `gorm:"uniqueIndex"`
	Applied bool
}

func InitialMigration(db *gorm.DB) error {
	// First create migrations table if it doesn't exist
	if err := db.AutoMigrate(&Migration{}); err != nil {
		return err
	}

	// Create your initial tables
	err := db.AutoMigrate(
		&models.User{},
		// Add other models here
	)
	if err != nil {
		return err
	}

	// Record this migration
	migration := Migration{
		Version: "000001",
		Applied: true,
	}

	return db.Create(&migration).Error
}
