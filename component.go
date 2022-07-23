package handler

import (
	"strings"

	"github.com/disgoorg/disgo/events"
)

type ComponentHandler func(args []string, e *events.ComponentInteractionCreate) error

type Component struct {
	Action  string
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

	if err := component.Handler(args[2:], e); err != nil {
		h.Logger.Errorf("Failed to handle component interaction for \"%s\" : %s", componentName, err)
	}
}
