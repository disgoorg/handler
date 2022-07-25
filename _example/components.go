package main

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/handler"
)

func TestComponent(b *Bot) handler.Component {
	return handler.Component{
		Name:  "test",
		Check: simpleComponentCheck(b),
		Handler: func(ctx *handler.ComponentContext) error {
			b.Logger.Info(ctx.Printer.Sprintf("components.test"))
			return ctx.CreateMessage(discord.MessageCreate{
				Content: ctx.Printer.Sprintf("checks.component"),
			})
		},
	}
}

func simpleComponentCheck(b *Bot) func(ctx *handler.ComponentContext) bool {
	return func(ctx *handler.ComponentContext) bool {
		b.Logger.Info(ctx.Printer.Sprintf("checks.component"))
		return ctx.User().ID == userID
	}
}
