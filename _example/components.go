package main

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/handler"
)

func TestComponent(b *Bot) handler.Component {
	return handler.Component{
		Name: "test",
		Handler: func(args []string, e *events.ComponentInteractionCreate) error {
			b.Logger.Info("test component")
			return e.CreateMessage(discord.MessageCreate{
				Content: "test button",
			})
		},
	}
}
