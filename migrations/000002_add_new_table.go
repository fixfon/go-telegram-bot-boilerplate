package migrations

import (
	"gorm.io/gorm"
)

func AddNewTable(db *gorm.DB) error {
	// Your migration code here
	// Example:
	type NewTable struct {
		gorm.Model
		Name string
	}

	return db.AutoMigrate(&NewTable{})
}
