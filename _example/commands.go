package main

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
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
			"test1": func(event *events.ApplicationCommandInteractionCreate) error {
				b.Logger.Info("Test Command 1")

				return event.CreateMessage(discord.MessageCreate{
					Content: "Test Command 1",
					Components: []discord.ContainerComponent{
						discord.ActionRowComponent{
							discord.NewPrimaryButton("test1", "handler:test"),
						},
					},
				})
			},
			"test/test2": func(event *events.ApplicationCommandInteractionCreate) error {
				b.Logger.Info("Test Command 2")

				return event.CreateModal(discord.ModalCreate{
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

func simpleCommandCheck(b *Bot) handler.Check[*events.ApplicationCommandInteractionCreate] {
	return func(event *events.ApplicationCommandInteractionCreate) bool {
		b.Logger.Info("Command Check")
		return event.User().ID == userID
	}
}
