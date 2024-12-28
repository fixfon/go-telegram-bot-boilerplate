package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	TelegramID int64  `gorm:"uniqueIndex"`
	Username   string `gorm:"size:255"`
	FirstName  string `gorm:"size:255"`
	LastName   string `gorm:"size:255"`
}
