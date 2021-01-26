package handlers

import (
	"atlas-clc/packets/inputs"
	"atlas-clc/packets/inputs/models"
	"atlas-clc/packets/inputs/readers"
	"atlas-clc/packets/outputs/writers"
	"atlas-clc/processors"
	"atlas-clc/registries"
	"atlas-clc/sessions"
	"log"
)

type CheckCharacterNameHandler struct {
}

func (h *CheckCharacterNameHandler) Handle(l *log.Logger, sessionId int, r *inputs.Reader) {
	s := registries.GetSessionRegistry().GetSession(sessionId)
	if s != nil {
		p := readers.ReadCheckCharacterName(r)
		if p != nil {
			h.handle(l, s, p)
		}
	}
}

func (h *CheckCharacterNameHandler) handle(l *log.Logger, s *sessions.Session, p *models.CheckCharacterName) {
	v, err := processors.IsValidName(l, p.Name())
	if err != nil {
		l.Println("[ERROR] validating character name on creation")
		s.Announce(writers.WriteCharacterNameCheckResponse(p.Name(), true))
	}
	s.Announce(writers.WriteCharacterNameCheckResponse(p.Name(), !v))
}
