package handler

import (
	"atlas-clc/character"
	"atlas-clc/session"
	"atlas-clc/socket/response/writer"
	"atlas-clc/world"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const OpCodeCharacterListAll uint16 = 0x0D

func HandleCharacterListAllRequest(l logrus.FieldLogger, span opentracing.Span) func(s session.Model, _ *request.RequestReader) {
	return func(s session.Model, _ *request.RequestReader) {
		ws, err := world.GetAll(l, span)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve worlds")
		}

		cm := getWorldCharacters(l, span)(s.AccountId(), ws)
		announceAllCharacters(l, cm, s)
	}
}

func announceAllCharacters(l logrus.FieldLogger, cm map[byte][]character.Model, ms session.Model) {
	cs := uint32(len(cm))
	unk := cs + (3 - cs%3) // row size

	err := session.Announce(writer.WriteShowAllCharacter(l)(cs, unk))(ms)
	if err != nil {
		l.WithError(err).Errorf("Unable to show all characters")
	}
	for k, v := range cm {
		err = session.Announce(writer.WriteShowAllCharacterInfo(l)(k, v, false))(ms)
		if err != nil {
			l.WithError(err).Errorf("Unable to show character information")
		}
	}
}

func getWorldCharacters(l logrus.FieldLogger, span opentracing.Span) func(accountId uint32, ws []world.Model) map[byte][]character.Model {
	return func(accountId uint32, ws []world.Model) map[byte][]character.Model {
		var cwm = make(map[byte][]character.Model, 0)
		for _, x := range ws {
			cs, err := character.GetForWorld(l, span)(accountId, x.Id())
			if err == nil {
				cwm[x.Id()] = cs
			}
		}
		return cwm
	}
}
