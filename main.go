package main

import (
	"fixfon/go-telegram-bot-boilerplate/config"
	"log"
	"time"

	client "gopkg.in/telebot.v4"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatal("Cannot load config:", err)
	}

	config.ConnectDB()

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
		return c.Send("Hello!")
	})

	bot.Start()
}
