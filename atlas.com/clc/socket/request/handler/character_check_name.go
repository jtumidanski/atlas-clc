package handler

import (
	"atlas-clc/character"
	"atlas-clc/session"
	"atlas-clc/socket/response/writer"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const OpCodeCharacterCheckName uint16 = 0x15

type CharacterCheckNameRequest struct {
	name string
}

func (c *CharacterCheckNameRequest) Name() string {
	return c.name
}

func ReadCharacterCheckNameRequest(reader *request.RequestReader) *CharacterCheckNameRequest {
	name := reader.ReadAsciiString()
	return &CharacterCheckNameRequest{name}
}

func HandleCheckCharacterNameRequest(l logrus.FieldLogger, span opentracing.Span) func(s session.Model, r *request.RequestReader) {
	return func(s session.Model, r *request.RequestReader) {
		p := ReadCharacterCheckNameRequest(r)

		ok, err := character.IsValidName(l, span)(p.Name())
		if err != nil {
			l.WithError(err).Errorf("Validating character name on creation")
			err = session.Announce(writer.WriteCharacterNameCheck(l)(p.Name(), true))(s)
			if err != nil {
				l.WithError(err).Errorf("Unable to issue character name validation error")
			}
		}
		err = session.Announce(writer.WriteCharacterNameCheck(l)(p.Name(), !ok))(s)
		if err != nil {
			l.WithError(err).Errorf("Unable to inform character name validation success")
		}
	}
}
