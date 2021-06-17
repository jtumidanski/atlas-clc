package handler

import (
	"atlas-clc/account"
	"atlas-clc/character"
	"atlas-clc/session"
	"atlas-clc/socket/response/writer"
	"atlas-clc/world"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/sirupsen/logrus"
)

const OpCodeCharacterListAll uint16 = 0x0D

type CharacterListAllHandler struct {
}

func (h *CharacterListAllHandler) IsValid(l logrus.FieldLogger, ms *session.MapleSession) bool {
	v := account.IsLoggedIn((*ms).AccountId())
	if !v {
		l.Errorf("Attempting to process a [CharacterListAlLRequest] when the account %d is not logged in.", (*ms).SessionId())
	}
	return v
}

func (h *CharacterListAllHandler) HandleRequest(l logrus.FieldLogger, ms *session.MapleSession, _ *request.RequestReader) {
	ws, err := world.GetWorlds()
	if err != nil {
		l.WithError(err).Errorf("Unable to retrieve worlds")
	}

	cm := h.getWorldCharacters((*ms).AccountId(), ws)
	h.announceAllCharacters(cm, ms)
}

func (h *CharacterListAllHandler) announceAllCharacters(cm map[byte][]character.Model, ms *session.MapleSession) {
	cs := uint32(len(cm))
	unk := cs + (3 - cs%3) // row size

	(*ms).Announce(writer.WriteShowAllCharacter(cs, unk))
	for k, v := range cm {
		(*ms).Announce(writer.WriteShowAllCharacterInfo(k, v, false))
	}
}

func (h *CharacterListAllHandler) getWorldCharacters(accountId uint32, ws []world.Model) map[byte][]character.Model {
	var cwm = make(map[byte][]character.Model, 0)
	for _, x := range ws {
		cs, err := character.GetCharactersForWorld(accountId, x.Id())
		if err == nil {
			cwm[x.Id()] = cs
		}
	}
	return cwm
}
