package handlers

import (
	"atlas-clc/packets/inputs"
	"atlas-clc/packets/inputs/constants"
	"log"
)

type Handler interface {
	Handle(l *log.Logger, sessionId int, r *inputs.Reader)
}

type Handle struct {
	l *log.Logger

	h Handler
}

func (h *Handle) Handle(sessionId int, r *inputs.Reader) {
	h.h.Handle(h.l, sessionId, r)
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
