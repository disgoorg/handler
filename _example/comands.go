package main

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/handler"
)

func TestCommand(b *Bot) handler.Command {
	return handler.Command{
		Create: discord.SlashCommandCreate{
			CommandName: "test",
			Description: "Test command",
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionSubCommand{
					CommandName: "test1",
					Description: "Test command 1",
				},
				discord.ApplicationCommandOptionSubCommandGroup{
					GroupName:   "test",
					Description: "Test command 1",
					Options: []discord.ApplicationCommandOptionSubCommand{
						{
							CommandName: "test2",
							Description: "Test command 2",
						},
					},
				},
			},
		},
		CommandHandlers: map[string]handler.CommandHandler{
			"test1": func(e *events.ApplicationCommandInteractionCreate) error {
				b.Logger.Info("test1")
				return nil
			},
			"test/test2": func(e *events.ApplicationCommandInteractionCreate) error {
				b.Logger.Info("test2")
				return nil
			},
		},
	}
}
