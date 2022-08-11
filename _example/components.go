package main

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/handler"
)

func TestComponent(b *Bot) handler.Component {
	return handler.Component{
		Name:  "test",
		Check: simpleComponentCheck(b),
		Handler: func(event *events.ComponentInteractionCreate) error {
			b.Logger.Info("Test Component")
			return event.CreateMessage(discord.MessageCreate{
				Content: "Test Component",
			})
		},
	}
}

func simpleComponentCheck(b *Bot) handler.Check[*events.ComponentInteractionCreate] {
	return func(event *events.ComponentInteractionCreate) bool {
		b.Logger.Info("Test Component Check")
		return event.User().ID == userID
	}
}
