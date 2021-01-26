package handlers

import (
	"atlas-clc/models"
	"atlas-clc/packets/inputs"
	"atlas-clc/packets/outputs/writers"
	"atlas-clc/processors"
	"atlas-clc/registries"
	"atlas-clc/sessions"
	"log"
)

type ViewAllCharactersHandler struct {
}

func (h *ViewAllCharactersHandler) Handle(l *log.Logger, sessionId int, _ *inputs.Reader) {
	s := registries.GetSessionRegistry().GetSession(sessionId)
	if s != nil {
		h.handle(l, s)
	}
}

func (h *ViewAllCharactersHandler) handle(l *log.Logger, s *sessions.Session) {
	ws, err := processors.GetWorlds(l)
	if err != nil {
		l.Println("[ERROR] unable to retrieve worlds")
	}

	cm := h.getWorldCharacters(l, s.AccountId(), ws)
	h.announceAllCharacters(cm, s)
}

func (h *ViewAllCharactersHandler) announceAllCharacters(cm map[byte][]models.Character, s *sessions.Session) {
	cs := uint32(len(cm))
	unk := cs + (3 - cs%3) // row size

	s.Announce(writers.WriteShowAllCharacter(cs, unk))
	for k, v := range cm {
		s.Announce(writers.WriteShowAllCharacterInfo(k, v, false))
	}
}

func (h *ViewAllCharactersHandler) getWorldCharacters(l *log.Logger, accountId int, ws []models.World) map[byte][]models.Character {
	var cwm = make(map[byte][]models.Character, 0)
	for _, x := range ws {
		cs, err := processors.GetCharactersForWorld(l, accountId, x.Id())
		if err == nil {
			cwm[x.Id()] = cs
		}
	}
	return cwm
}
