package handler

import (
	"atlas-clc/sessions"
	"atlas-clc/socket/request"
	"log"
)

const OpCodePong uint16 = 0x18

type PongHandler struct {
}

func (h *PongHandler) IsValid(_ *log.Logger, _ *sessions.Session) bool {
	return true
}

func (h *PongHandler) HandleRequest(_ *log.Logger, _ *sessions.Session, _ *request.RequestReader) {
}
