package main

import (
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/handler"
)

func TestModal(b *Bot) handler.Modal {
	return handler.Modal{
		Action: "test",
		Handler: func(args []string, e *events.ModalSubmitInteractionCreate) error {
			b.Logger.Info("test modal")
			return nil
		},
	}
}
