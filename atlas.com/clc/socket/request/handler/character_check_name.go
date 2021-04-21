package handler

import (
	"atlas-clc/mapleSession"
	"atlas-clc/processors"
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

type CharacterCheckNameHandler struct {
}

func (h *CharacterCheckNameHandler) IsValid(l logrus.FieldLogger, ms *mapleSession.MapleSession) bool {
	v := processors.IsLoggedIn((*ms).AccountId())
	if !v {
		l.Errorf("Attempting to process a [CharacterCheckNameRequest] when the account %d is not logged in.", (*ms).SessionId())
	}
	return v
}

func (h *CharacterCheckNameHandler) HandleRequest(l logrus.FieldLogger, ms *mapleSession.MapleSession, r *request.RequestReader) {
	p := ReadCharacterCheckNameRequest(r)

	v, err := processors.IsValidName(p.Name())
	if err != nil {
		l.WithError(err).Errorf("Validating character name on creation")
		(*ms).Announce(writer.WriteCharacterNameCheck(p.Name(), true))
	}
	(*ms).Announce(writer.WriteCharacterNameCheck(p.Name(), !v))
}
