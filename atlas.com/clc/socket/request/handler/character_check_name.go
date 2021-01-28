package handler

import (
	"atlas-clc/processors"
	"atlas-clc/sessions"
	"atlas-clc/socket/request"
	"atlas-clc/socket/response/writer"
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

func (h *CharacterCheckNameHandler) IsValid(l *log.Logger, s *sessions.Session) bool {
	v := processors.IsLoggedIn(s.AccountId())
	if !v {
		l.Printf("[ERROR] attempting to process a [CharacterCheckNameRequest] when the account %d is not logged in.", s.SessionId())
	}
	return v
}

func (h *CharacterCheckNameHandler) HandleRequest(l *log.Logger, s *sessions.Session, r *request.RequestReader) {
	p := ReadCharacterCheckNameRequest(r)

	v, err := processors.IsValidName(p.Name())
	if err != nil {
		l.Println("[ERROR] validating character name on creation")
		s.Announce(writer.WriteCharacterNameCheckResponse(p.Name(), true))
	}
	s.Announce(writer.WriteCharacterNameCheckResponse(p.Name(), !v))
}
