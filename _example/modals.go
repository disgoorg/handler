package main

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/handler"
)

func TestModal(b *Bot) handler.Modal {
	return handler.Modal{
		Name:  "test",
		Check: simpleModalCheck(b),
		Handler: func(event *events.ModalSubmitInteractionCreate) error {
			b.Logger.Info("Test Modal")
			return event.CreateMessage(discord.MessageCreate{
				Content: fmt.Sprintf("test modal: %s", event.Data.Text("test-input")),
			})
		},
	}
}

func simpleModalCheck(b *Bot) handler.Check[*events.ModalSubmitInteractionCreate] {
	return func(event *events.ModalSubmitInteractionCreate) bool {
		b.Logger.Info("Modal Check")
		return event.User().ID == userID
	}
}
