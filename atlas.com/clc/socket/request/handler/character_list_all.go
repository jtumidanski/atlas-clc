package handler

import (
	"atlas-clc/character"
	"atlas-clc/session"
	"atlas-clc/socket/response/writer"
	"atlas-clc/world"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/sirupsen/logrus"
)

const OpCodeCharacterListAll uint16 = 0x0D

func HandleCharacterListAllRequest(l logrus.FieldLogger, ms *session.Model, _ *request.RequestReader) {
	ws, err := world.GetAll(l)
	if err != nil {
		l.WithError(err).Errorf("Unable to retrieve worlds")
	}

	cm := getWorldCharacters(l)(ms.AccountId(), ws)
	announceAllCharacters(l, cm, ms)
}

func announceAllCharacters(l logrus.FieldLogger, cm map[byte][]character.Model, ms *session.Model) {
	cs := uint32(len(cm))
	unk := cs + (3 - cs%3) // row size

	err := ms.Announce(writer.WriteShowAllCharacter(l)(cs, unk))
	if err != nil {
		l.WithError(err).Errorf("Unable to show all characters")
	}
	for k, v := range cm {
		err = ms.Announce(writer.WriteShowAllCharacterInfo(l)(k, v, false))
		if err != nil {
			l.WithError(err).Errorf("Unable to show character information")
		}
	}
}

func getWorldCharacters(l logrus.FieldLogger) func(accountId uint32, ws []world.Model) map[byte][]character.Model {
	return func(accountId uint32, ws []world.Model) map[byte][]character.Model {
		var cwm = make(map[byte][]character.Model, 0)
		for _, x := range ws {
			cs, err := character.GetForWorld(l)(accountId, x.Id())
			if err == nil {
				cwm[x.Id()] = cs
			}
		}
		return cwm
	}
}
