package main

import (
	"fixfon/go-telegram-bot-boilerplate/commands"
	"fixfon/go-telegram-bot-boilerplate/config"
	"fixfon/go-telegram-bot-boilerplate/migrations"
	"log"
	"time"

	client "gopkg.in/telebot.v4"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// Connect to database
	if err := config.ConnectDB(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	// Run migrations after database connection is established
	if err := migrations.RunMigrations(config.GetDB()); err != nil {
		log.Fatal("Cannot run migrations:", err)
	}

	pref := client.Settings{
		Token:  config.AppConfig.TelegramToken,
		Poller: &client.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := client.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle("/register", func(c client.Context) error {
		return commands.Register(c)
	})

	bot.Start()
}
