package commands

import (
	client "gopkg.in/telebot.v4"
)

func Register(c client.Context) error {
	// check if the sender is admin
	if c.Sender().ID != 123456 {
		return c.Send("You are not authorized to use this command")
	}

	// register user to the registered users list
	
}
