package handler

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
)

var _ bot.EventListener = (*Handler)(nil)

func New(logger log.Logger) *Handler {
	return &Handler{
		Logger: logger,
	}
}

type Handler struct {
	Logger log.Logger

	Commands   map[string]Command
	Components map[string]Component
	Modals     map[string]Modal
}

func (h *Handler) AddCommands(commands ...Command) {
	for _, command := range commands {
		h.Commands[command.Create.Name()] = command
	}
}

func (h *Handler) AddComponents(components ...Component) {
	for _, component := range components {
		h.Components[component.Action] = component
	}
}

func (h *Handler) AddModals(modals ...Modal) {
	for _, modal := range modals {
		h.Modals[modal.Action] = modal
	}
}

func (h *Handler) SyncCommands(client bot.Client, guildIDs ...snowflake.ID) {
	commands := make([]discord.ApplicationCommandCreate, len(h.Commands))
	for _, command := range h.Commands {
		commands = append(commands, command.Create)
	}

	if len(guildIDs) == 0 {
		if _, err := client.Rest().SetGlobalCommands(client.ApplicationID(), commands); err != nil {
			h.Logger.Error("Failed to sync global commands: ", err)
			return
		}
		h.Logger.Infof("Synced %d global commands", len(commands))
		return
	}

	for _, guildID := range guildIDs {
		if _, err := client.Rest().SetGuildCommands(client.ApplicationID(), guildID, commands); err != nil {
			h.Logger.Errorf("Failed to sync commands for guild %d: %s", guildID, err)
			continue
		}
		h.Logger.Infof("Synced %d commands for guild %s", len(commands), guildID)
	}
}

func (h *Handler) OnEvent(event bot.Event) {
	switch e := event.(type) {
	case *events.ApplicationCommandInteractionCreate:
		h.handleCommand(e)
	case *events.AutocompleteInteractionCreate:
		h.handleAutocomplete(e)
	case *events.ComponentInteractionCreate:
		h.handleComponent(e)
	case *events.ModalSubmitInteractionCreate:
		h.handleModal(e)
	}
}
