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
	}
	return nil
}
