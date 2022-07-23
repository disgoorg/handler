package handler

import (
	"strings"

	"github.com/disgoorg/disgo/events"
)

type ModalHandler func(args []string, e *events.ModalSubmitInteractionCreate) error

type Modal struct {
	Name    string
	Handler ModalHandler
}

func (h *Handler) handleModal(e *events.ModalSubmitInteractionCreate) {
	customID := e.Data.CustomID.String()
	if !strings.HasPrefix(customID, "handler:") {
		return
	}

	args := strings.Split(customID, ":")
	modalName := args[1]
	modal, ok := h.Modals[modalName]
	if !ok || modal.Handler == nil {
		h.Logger.Errorf("No modal handler for \"%s\" found", modalName)
	}

	if err := modal.Handler(args[2:], e); err != nil {
		h.Logger.Errorf("Failed to handle modal interaction for \"%s\" : %s", modalName, err)
	}
}
