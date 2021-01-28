package handlers

import (
	"atlas-clc/packets/inputs"
	"atlas-clc/packets/inputs/constants"
	"atlas-clc/registries"
	"atlas-clc/sessions"
	"log"
)

type Handler interface {
	IsValid(l *log.Logger, s *sessions.Session) bool

	Handle(l *log.Logger, s *sessions.Session, r *inputs.Reader)
}

type Handle struct {
	l *log.Logger

	h Handler
}

func (h *Handle) Handle(sessionId int, r *inputs.Reader) {
	s := registries.GetSessionRegistry().GetSession(sessionId)
	if s != nil {
		if h.h.IsValid(h.l, s) {
			h.h.Handle(h.l, s, r)
		}
	} else {
		h.l.Printf("[ERROR] unable to locate session %d", sessionId)
	}
}

func GetHandle(l *log.Logger, op uint16) *Handle {
	switch op {
	case constants.LoginPassword:
		return &Handle{l, &LoginPasswordHandler{}}
	case constants.ServerListReRequest:
		return &Handle{l, &ServerListRequestHandler{}}
	case constants.CharacterListRequest:
		return &Handle{l, &CharacterListRequestHandler{}}
	case constants.ServerStatusRequest:
		return &Handle{l, &ServerStatusHandler{}}
	case constants.ServerListRequest:
		return &Handle{l, &ServerListRequestHandler{}}
	case constants.ViewAllCharacters:
		return &Handle{l, &ViewAllCharactersHandler{}}
	case constants.PickAllCharacters:
		return &Handle{l, &PickAllCharactersHandler{}}
	case constants.CharacterSelect:
		return &Handle{l, &CharacterSelectHandler{}}
	case constants.CheckCharacterName:
		return &Handle{l, &CheckCharacterNameHandler{}}
	case constants.CreateCharacter:
		return &Handle{l, &CreateCharacterHandler{}}
	case constants.Pong:
		return &Handle{l, &PongHandler{}}
	}
	return nil
}
