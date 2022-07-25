package main

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/handler"
)

func TestCommand(b *Bot) handler.Command {
	return handler.Command{
		Create: discord.SlashCommandCreate{
			Name:        "test",
			Description: "Test command",
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionSubCommand{
					Name:        "test1",
					Description: "Test command 1",
				},
				discord.ApplicationCommandOptionSubCommandGroup{
					Name:        "test",
					Description: "Test command 1",
					Options: []discord.ApplicationCommandOptionSubCommand{
						{
							Name:        "test2",
							Description: "Test command 2",
						},
					},
				},
			},
		},
		Check: simpleCommandCheck(b),
		CommandHandlers: map[string]handler.CommandHandler{
			"test1": func(ctx *handler.CommandContext) error {
				b.Logger.Info(ctx.Printer.Sprintf("commands.test1"))

				return ctx.CreateMessage(discord.MessageCreate{
					Content: ctx.Printer.Sprintf("commands.test1"),
					Components: []discord.ContainerComponent{
						discord.ActionRowComponent{
							discord.NewPrimaryButton("test1", "handler:test"),
						},
					},
				})
			},
			"test/test2": func(ctx *handler.CommandContext) error {
				b.Logger.Info(ctx.Printer.Sprintf("commands.test2"))

				return ctx.CreateModal(discord.ModalCreate{
					CustomID: "handler:test",
					Title:    "Test Modal",
					Components: []discord.ContainerComponent{
						discord.ActionRowComponent{
							discord.NewShortTextInput("test-input", "test input"),
						},
					},
				})
			},
		},
	}
}

func simpleCommandCheck(b *Bot) func(ctx *handler.CommandContext) bool {
	return func(ctx *handler.CommandContext) bool {
		b.Logger.Info(ctx.Printer.Sprintf("checks.command"))
		return ctx.User().ID == userID
	}
}
