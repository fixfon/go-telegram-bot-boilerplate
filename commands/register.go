package commands

import (
	"log"

	"fixfon/go-telegram-bot-boilerplate/config"
	"fixfon/go-telegram-bot-boilerplate/models"
	"fixfon/go-telegram-bot-boilerplate/repository"

	client "gopkg.in/telebot.v4"
)

func Register(c client.Context) error {
	senderID := c.Sender().ID
	messageContent := c.Message().Text
	log.Printf("Sender ID: %d, Message Content: %s", senderID, messageContent)

	// Initialize repository
	userRepo := repository.NewUserRepository(config.DB)

	// Check if user is already registered
	if userRepo.IsUserRegistered(senderID) {
		return c.Send("This user is already registered!")
	}

	// Check if sender is owner
	if senderID != config.AppConfig.OwnerID {
		return c.Send("Sorry, only the owner can register!")
	}

	// Create new user
	user := &models.User{
		TelegramID: senderID,
		Username:   c.Sender().Username,
		FirstName:  c.Sender().FirstName,
		LastName:   c.Sender().LastName,
		IsOwner:    false,
	}

	if err := userRepo.Create(user); err != nil {
		log.Printf("Error creating user: %v", err)
		return c.Send("Failed to register. Please try again later.")
	}

	return c.Send("You have been registered by the owner!")
}
