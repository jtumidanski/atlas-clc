package handler

import (
	"atlas-clc/domain"
	"atlas-clc/processors"
	"atlas-clc/sessions"
	"atlas-clc/socket/request"
	"atlas-clc/socket/response/writer"
	"log"
)

const OpCodeCharacterListAll uint16 = 0x0D

type CharacterListAllHandler struct {
}

func (h *CharacterListAllHandler) IsValid(l *log.Logger, s *sessions.Session) bool {
	v := processors.IsLoggedIn(s.AccountId())
	if !v {
		l.Printf("[ERROR] attempting to process a [CharacterListAlLRequest] when the account %d is not logged in.", s.SessionId())
	}
	return v
}

func (h *CharacterListAllHandler) HandleRequest(l *log.Logger, s *sessions.Session, _ *request.RequestReader) {
	ws, err := processors.GetWorlds()
	if err != nil {
		l.Println("[ERROR] unable to retrieve worlds")
	}

	cm := h.getWorldCharacters(l, s.AccountId(), ws)
	h.announceAllCharacters(cm, s)
}

func (h *CharacterListAllHandler) announceAllCharacters(cm map[byte][]domain.Character, s *sessions.Session) {
	cs := uint32(len(cm))
	unk := cs + (3 - cs%3) // row size

	s.Announce(writer.WriteShowAllCharacter(cs, unk))
	for k, v := range cm {
		s.Announce(writer.WriteShowAllCharacterInfo(k, v, false))
	}
}

func (h *CharacterListAllHandler) getWorldCharacters(l *log.Logger, accountId int, ws []domain.World) map[byte][]domain.Character {
	var cwm = make(map[byte][]domain.Character, 0)
	for _, x := range ws {
		cs, err := processors.GetCharactersForWorld(accountId, x.Id())
		if err == nil {
			cwm[x.Id()] = cs
		}
	}
	return cwm
}
