package handlers

import (
	models2 "atlas-clc/models"
	"atlas-clc/packets/inputs"
	"atlas-clc/packets/inputs/models"
	"atlas-clc/packets/inputs/readers"
	"atlas-clc/packets/outputs/writers"
	"atlas-clc/processors"
	"atlas-clc/registries"
	"atlas-clc/sessions"
	"log"
)

type CharacterListRequestHandler struct {
}

func (h *CharacterListRequestHandler) Handle(l *log.Logger, sessionId int, r *inputs.Reader) {
	s := registries.GetSessionRegistry().GetSession(sessionId)
	if s != nil {
		p := readers.ReadCharacterListRequest(r)
		if p != nil {
			h.handle(l, s, p)
		}
	}
}

func (h *CharacterListRequestHandler) handle(l *log.Logger, s *sessions.Session, p *models.CharacterListRequest) {
	w, err := processors.GetWorld(l, p.WorldId())
	if err != nil {
		l.Println("[ERROR] received a character list request for a world we do not have")
		return
	}

	if w.CapacityStatus() == models2.Full {
		s.Announce(writers.WriteWorldCapacityStatus(models2.Full))
		return
	}

	s.SetWorldId(p.WorldId())
	s.SetChannelId(p.ChannelId())

	a, err := processors.GetAccountById(l, s.AccountId())
	if err != nil {
		l.Println("[ERROR] cannot retrieve account")
		return
	}

	cs, err := processors.GetCharactersForWorld(l, s.AccountId(), p.WorldId())
	if err != nil {
		l.Println("[ERROR] cannot retrieve account characters")
		return
	}

	s.Announce(writers.WriteCharacterList(cs, p.WorldId(), 0, true, a.PIC(), int16(1), a.CharacterSlots()))
}
