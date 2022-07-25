package handler

import (
	"strings"

	"github.com/disgoorg/disgo/events"
)

type ModalHandler func(ctx *ModalContext) error

type Modal struct {
	Name    string
	Check   Check[*ModalContext]
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

	ctx := &ModalContext{
		ModalSubmitInteractionCreate: e,
		Printer:                      h.I18n.NewPrinter(e.Locale()),
		Args:                         args[2:],
	}

	if modal.Check != nil && !modal.Check(ctx) {
		return
	}

	if err := modal.Handler(ctx); err != nil {
		h.Logger.Errorf("Failed to handle modal interaction for \"%s\" : %s", modalName, err)
	}
}
