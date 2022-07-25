package handler

import (
	"strings"

	"github.com/disgoorg/disgo/events"
)

type ComponentHandler func(ctx *ComponentContext) error

type Component struct {
	Name    string
	Check   Check[*ComponentContext]
	Handler ComponentHandler
}

func (h *Handler) handleComponent(e *events.ComponentInteractionCreate) {
	customID := e.Data.CustomID().String()
	if !strings.HasPrefix(customID, "handler:") {
		return
	}

	args := strings.Split(customID, ":")
	componentName := args[1]
	component, ok := h.Components[componentName]
	if !ok || component.Handler == nil {
		h.Logger.Errorf("No component handler for \"%s\" found", componentName)
	}

	ctx := &ComponentContext{
		ComponentInteractionCreate: e,
		Printer:                    h.I18n.NewPrinter(e.Locale()),
		Args:                       args[2:],
	}

	if component.Check != nil && !component.Check(ctx) {
		return
	}

	if err := component.Handler(ctx); err != nil {
		h.Logger.Errorf("Failed to handle component interaction for \"%s\" : %s", componentName, err)
	}
}
