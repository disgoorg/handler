package main

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/handler"
)

func TestModal(b *Bot) handler.Modal {
	return handler.Modal{
		Name: "test",
		Handler: func(args []string, e *events.ModalSubmitInteractionCreate) error {
			b.Logger.Info("test modal")
			return e.CreateMessage(discord.MessageCreate{
				Content: fmt.Sprintf("test modal: %s", e.Data.Text("test-input")),
			})
		},
	}
}
