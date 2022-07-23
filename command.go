package handler

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

type (
	CommandHandler      func(e *events.ApplicationCommandInteractionCreate) error
	AutocompleteHandler func(e *events.AutocompleteInteractionCreate) error
)

type Command struct {
	Create               discord.ApplicationCommandCreate
	CommandHandlers      map[string]CommandHandler
	AutocompleteHandlers map[string]AutocompleteHandler
}

func (h *Handler) handleCommand(e *events.ApplicationCommandInteractionCreate) {
	name := e.Data.CommandName()
	cmd, ok := h.Commands[name]
	if !ok || cmd.CommandHandlers == nil {
		h.Logger.Errorf("No command or handler found for \"%s\"", name)
	}

	var path string
	if d, ok := e.Data.(discord.SlashCommandInteractionData); ok {
		path = buildCommandPath(d.SubCommandName, d.SubCommandGroupName)
	}

	handler, ok := cmd.CommandHandlers[path]
	if !ok {
		h.Logger.Warnf("No handler for command \"%s\" with path \"%s\" found", name, path)
		return
	}

	if err := handler(e); err != nil {
		h.Logger.Errorf("Failed to handle command \"%s\" with path \"%s\": %s", name, path, err)
	}
}

func (h *Handler) handleAutocomplete(e *events.AutocompleteInteractionCreate) {
	name := e.Data.CommandName
	cmd, ok := h.Commands[name]
	if !ok || cmd.AutocompleteHandlers == nil {
		h.Logger.Errorf("No command or handler found for \"%s\"", name)
	}

	path := buildCommandPath(e.Data.SubCommandName, e.Data.SubCommandGroupName)

	handler, ok := cmd.AutocompleteHandlers[path]
	if !ok {
		h.Logger.Warnf("No autocomplete handler for command \"%s\" with path \"%s\" found", name, path)
		return
	}

	if err := handler(e); err != nil {
		h.Logger.Errorf("Failed to handle autocomplete for command \"%s\" with path \"%s\": %s", name, path, err)
	}
}

func buildCommandPath(subcommand *string, subcommandGroup *string) string {
	var path string
	if subcommand != nil {
		path = *subcommand
	}
	if subcommandGroup != nil {
		path = *subcommandGroup + "/" + path
	}
	return path
}
