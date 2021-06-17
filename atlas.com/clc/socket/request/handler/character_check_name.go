package handler

import (
	"atlas-clc/character"
	"atlas-clc/session"
	"atlas-clc/socket/response/writer"
	"github.com/jtumidanski/atlas-socket/request"
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

func HandleCheckCharacterNameRequest(l logrus.FieldLogger, ms *session.MapleSession, r *request.RequestReader) {
	p := ReadCharacterCheckNameRequest(r)

	ok, err := character.IsValidName(p.Name())
	if err != nil {
		l.WithError(err).Errorf("Validating character name on creation")
		err = (*ms).Announce(writer.WriteCharacterNameCheck(l)(p.Name(), true))
		if err != nil {
			l.WithError(err).Errorf("Unable to issue character name validation error")
		}
	}
	err = (*ms).Announce(writer.WriteCharacterNameCheck(l)(p.Name(), !ok))
	if err != nil {
		l.WithError(err).Errorf("Unable to inform character name validation success")
	}
}
