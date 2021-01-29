package handler

import (
	"atlas-clc/mapleSession"
	"atlas-clc/processors"
	"atlas-clc/socket/response/writer"
	"github.com/jtumidanski/atlas-socket/request"
	"log"
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

func (h *CharacterCheckNameHandler) IsValid(l *log.Logger, ms *mapleSession.MapleSession) bool {
	v := processors.IsLoggedIn((*ms).AccountId())
	if !v {
		l.Printf("[ERROR] attempting to process a [CharacterCheckNameRequest] when the account %d is not logged in.", (*ms).SessionId())
	}
	return v
}

func (h *CharacterCheckNameHandler) HandleRequest(l *log.Logger, ms *mapleSession.MapleSession, r *request.RequestReader) {
	p := ReadCharacterCheckNameRequest(r)

	v, err := processors.IsValidName(p.Name())
	if err != nil {
		l.Println("[ERROR] validating character name on creation")
		(*ms).Announce(writer.WriteCharacterNameCheckResponse(p.Name(), true))
	}
	(*ms).Announce(writer.WriteCharacterNameCheckResponse(p.Name(), !v))
}
