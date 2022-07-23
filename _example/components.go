package main

import (
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/handler"
)

func TestComponent(b *Bot) handler.Component {
	return handler.Component{
		Action: "test",
		Handler: func(args []string, e *events.ComponentInteractionCreate) error {
			b.Logger.Info("test component")
			return nil
		},
	}
}
