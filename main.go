package main

import (
	"log"
	"time"

	client "gopkg.in/telebot.v4"
)

func main() {
	pref := client.Settings{
		// Token:  os.Getenv("TOKEN"),
		Token:  "7616059022:AAEoSX9QZxPKmH2WcdRWNkZZMpkfXkVjywQ",
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
