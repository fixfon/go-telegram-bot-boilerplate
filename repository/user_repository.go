package repository

import (
	"fixfon/go-telegram-bot-boilerplate/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByTelegramID(telegramID int64) (*models.User, error) {
	var user models.User
	err := r.db.Where("telegram_id = ?", telegramID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) IsUserRegistered(telegramID int64) bool {
	var count int64
	r.db.Model(&models.User{}).Where("telegram_id = ?", telegramID).Count(&count)
	return count > 0
}
