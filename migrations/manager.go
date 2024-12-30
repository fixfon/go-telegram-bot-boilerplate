package migrations

import (
	"fmt"

	"gorm.io/gorm"
)

var migrations = []struct {
	version string
	up      func(*gorm.DB) error
}{
	{
		version: "000001",
		up:      InitialMigration,
	},
	// {
	// 	version: "000002",
	// 	up:      AddNewTable,
	// },
	// Add new migrations here
}

func RunMigrations(db *gorm.DB) error {

	// Check and run each migration
	for _, migration := range migrations {
		var existing Migration
		if err := db.Where("version = ?", migration.version).First(&existing).Error; err == nil {
			// Migration already applied
			fmt.Printf("Migration %s already applied\n", migration.version)
			continue
		}

		// Run migration
		fmt.Printf("Running migration %s...\n", migration.version)
		if err := migration.up(db); err != nil {
			return fmt.Errorf("error running migration %s: %v", migration.version, err)
		}

		// Record successful migration
		record := Migration{
			Version: migration.version,
			Applied: true,
		}
		if err := db.Create(&record).Error; err != nil {
			return fmt.Errorf("error recording migration %s: %v", migration.version, err)
		}

		fmt.Printf("Migration %s completed successfully\n", migration.version)
	}

	return nil
}
