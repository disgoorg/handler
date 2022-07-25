package handler

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

type (
	CommandHandler      func(ctx *CommandContext) error
	AutocompleteHandler func(ctx *AutocompleteContext) error
)

type Command struct {
	Create               discord.ApplicationCommandCreate
	Check                Check[*CommandContext]
	AutocompleteCheck    Check[*AutocompleteContext]
	CommandHandlers      map[string]CommandHandler
	AutocompleteHandlers map[string]AutocompleteHandler
}

func (h *Handler) handleCommand(e *events.ApplicationCommandInteractionCreate) {
	name := e.Data.CommandName()
	cmd, ok := h.Commands[name]
	if !ok || cmd.CommandHandlers == nil {
		h.Logger.Errorf("No command or handler found for \"%s\"", name)
	}

	ctx := &CommandContext{
		ApplicationCommandInteractionCreate: e,
		Printer:                             h.I18n.NewPrinter(e.Locale()),
	}

	if cmd.Check != nil && !cmd.Check(ctx) {
		return
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

	if err := handler(ctx); err != nil {
		h.Logger.Errorf("Failed to handle command \"%s\" with path \"%s\": %s", name, path, err)
	}
}

func (h *Handler) handleAutocomplete(e *events.AutocompleteInteractionCreate) {
	name := e.Data.CommandName
	cmd, ok := h.Commands[name]
	if !ok || cmd.AutocompleteHandlers == nil {
		h.Logger.Errorf("No command or handler found for \"%s\"", name)
	}

	ctx := &AutocompleteContext{
		AutocompleteInteractionCreate: e,
		Printer:                       h.I18n.NewPrinter(e.Locale()),
	}

	if cmd.AutocompleteCheck != nil && !cmd.AutocompleteCheck(ctx) {
		return
	}

	path := buildCommandPath(e.Data.SubCommandName, e.Data.SubCommandGroupName)

	handler, ok := cmd.AutocompleteHandlers[path]
	if !ok {
		h.Logger.Warnf("No autocomplete handler for command \"%s\" with path \"%s\" found", name, path)
		return
	}

	if err := handler(ctx); err != nil {
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
