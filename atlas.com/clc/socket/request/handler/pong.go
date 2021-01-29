package handler

import (
	"atlas-clc/mapleSession"
	"github.com/jtumidanski/atlas-socket/request"
	"log"
)

const OpCodePong uint16 = 0x18

type PongHandler struct {
}

func (h *PongHandler) IsValid(_ *log.Logger, _ *mapleSession.MapleSession) bool {
	return true
}

func (h *PongHandler) HandleRequest(_ *log.Logger, _ *mapleSession.MapleSession, _ *request.RequestReader) {
}
