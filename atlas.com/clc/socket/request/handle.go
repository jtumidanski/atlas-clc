package request

import (
	"atlas-clc/registries"
	"atlas-clc/sessions"
	"atlas-clc/socket/request/handler"
	"log"
)

type RequestHandler interface {
	IsValid(l *log.Logger, s *sessions.Session) bool

	HandleRequest(l *log.Logger, s *sessions.Session, r *RequestReader)
}

type RequestHandle struct {
	l *log.Logger

	h RequestHandler
}

func (h *RequestHandle) Handle(sessionId int, r *RequestReader) {
	s := registries.GetSessionRegistry().Get(sessionId)
	if s != nil {
		if h.h.IsValid(h.l, s) {
			h.h.HandleRequest(h.l, s, r)
			s.UpdateLastPacket()
		}
	} else {
		h.l.Printf("[ERROR] unable to locate session %d", sessionId)
	}
}

func GetHandle(l *log.Logger, op uint16) *RequestHandle {
	switch op {
	case handler.OpCodeLogin:
		return &RequestHandle{l, &handler.LoginHandler{}}
	case handler.OpCodeServerListReRequest:
		return &RequestHandle{l, &handler.ServerListHandler{}}
	case handler.OpCodeCharacterListWorld:
		return &RequestHandle{l, &handler.CharacterListWorldHandler{}}
	case handler.OpCodeServerStatus:
		return &RequestHandle{l, &handler.ServerStatusHandler{}}
	case handler.OpCodeServerRequest:
		return &RequestHandle{l, &handler.ServerListHandler{}}
	case handler.OpCodeCharacterListAll:
		return &RequestHandle{l, &handler.CharacterListAllHandler{}}
	case handler.OpCodeCharacterSelectFromAll:
		return &RequestHandle{l, &handler.CharacterSelectFromAllHandler{}}
	case handler.OpCodeCharacterSelectFromWorld:
		return &RequestHandle{l, &handler.CharacterSelectFromWorldHandler{}}
	case handler.OpCodeCharacterCheckName:
		return &RequestHandle{l, &handler.CharacterCheckNameHandler{}}
	case handler.OpCodeCharacterCreate:
		return &RequestHandle{l, &handler.CharacterCreateHandler{}}
	case handler.OpCodePong:
		return &RequestHandle{l, &handler.PongHandler{}}
	}
	return nil
}
