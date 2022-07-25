package handler

import (
	"github.com/disgoorg/disgo/events"
	"golang.org/x/text/message"
)

type CommandContext struct {
	*events.ApplicationCommandInteractionCreate
	Printer *message.Printer
}

type AutocompleteContext struct {
	*events.AutocompleteInteractionCreate
	Printer *message.Printer
}

type ComponentContext struct {
	*events.ComponentInteractionCreate
	Printer *message.Printer
	Args    []string
}

type ModalContext struct {
	*events.ModalSubmitInteractionCreate
	Printer *message.Printer
	Args    []string
}
