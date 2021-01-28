package handlers

import (
	"atlas-clc/packets/inputs"
	"atlas-clc/sessions"
	"log"
)

type PongHandler struct {
}

func (h *PongHandler) IsValid(l *log.Logger, s *sessions.Session) bool {
	return true
}

func (h *PongHandler) Handle(_ *log.Logger, _ *sessions.Session, _ *inputs.Reader) {
}
