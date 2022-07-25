package main

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/handler"
)

func TestModal(b *Bot) handler.Modal {
	return handler.Modal{
		Name:  "test",
		Check: simpleModalCheck(b),
		Handler: func(ctx *handler.ModalContext) error {
			b.Logger.Info(ctx.Printer.Sprintf("modals.test"))
			return ctx.CreateMessage(discord.MessageCreate{
				Content: fmt.Sprintf("test modal: %s", ctx.Data.Text("test-input")),
			})
		},
	}
}

func simpleModalCheck(b *Bot) func(ctx *handler.ModalContext) bool {
	return func(ctx *handler.ModalContext) bool {
		b.Logger.Info(ctx.Printer.Sprintf("checks.modal"))
		return ctx.User().ID == userID
	}
}
