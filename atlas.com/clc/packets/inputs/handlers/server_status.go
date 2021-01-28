package handlers

import (
	"atlas-clc/packets/inputs"
	"atlas-clc/packets/inputs/models"
	"atlas-clc/packets/inputs/readers"
	"atlas-clc/packets/outputs/writers"
	"atlas-clc/processors"
	"atlas-clc/sessions"
	"log"
)

type ServerStatusHandler struct {
}

func (h *ServerStatusHandler) IsValid(l *log.Logger, s *sessions.Session) bool {
	return processors.IsLoggedIn(l, s.AccountId())
}

func (h *ServerStatusHandler) Handle(l *log.Logger, s *sessions.Session, r *inputs.Reader) {
	p := readers.ReadServerStatus(r)
	if p != nil {
		h.handle(l, s, p)
	}
}

func (h *ServerStatusHandler) handle(l *log.Logger, s *sessions.Session, p *models.ServerStatus) {
	cs := processors.GetWorldCapacityStatus(l, p.WorldId())
	s.Announce(writers.WriteWorldCapacityStatus(cs))
}
