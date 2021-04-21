package handler

import (
	"atlas-clc/domain"
	"atlas-clc/mapleSession"
	"atlas-clc/processors"
	"atlas-clc/socket/response/writer"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/sirupsen/logrus"
)

const OpCodeCharacterListAll uint16 = 0x0D

type CharacterListAllHandler struct {
}

func (h *CharacterListAllHandler) IsValid(l logrus.FieldLogger, ms *mapleSession.MapleSession) bool {
	v := processors.IsLoggedIn((*ms).AccountId())
	if !v {
		l.Errorf("Attempting to process a [CharacterListAlLRequest] when the account %d is not logged in.", (*ms).SessionId())
	}
	return v
}

func (h *CharacterListAllHandler) HandleRequest(l logrus.FieldLogger, ms *mapleSession.MapleSession, _ *request.RequestReader) {
	ws, err := processors.GetWorlds()
	if err != nil {
		l.WithError(err).Errorf("Unable to retrieve worlds")
	}

	cm := h.getWorldCharacters((*ms).AccountId(), ws)
	h.announceAllCharacters(cm, ms)
}

func (h *CharacterListAllHandler) announceAllCharacters(cm map[byte][]domain.Character, ms *mapleSession.MapleSession) {
	cs := uint32(len(cm))
	unk := cs + (3 - cs%3) // row size

	(*ms).Announce(writer.WriteShowAllCharacter(cs, unk))
	for k, v := range cm {
		(*ms).Announce(writer.WriteShowAllCharacterInfo(k, v, false))
	}
}

func (h *CharacterListAllHandler) getWorldCharacters(accountId uint32, ws []domain.World) map[byte][]domain.Character {
	var cwm = make(map[byte][]domain.Character, 0)
	for _, x := range ws {
		cs, err := processors.GetCharactersForWorld(accountId, x.Id())
		if err == nil {
			cwm[x.Id()] = cs
		}
	}
	return cwm
}
