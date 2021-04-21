package handler

import (
	"atlas-clc/mapleSession"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/sirupsen/logrus"
)

const OpCodePong uint16 = 0x18

type PongHandler struct {
}

func (h *PongHandler) IsValid(_ logrus.FieldLogger, _ *mapleSession.MapleSession) bool {
	return true
}

func (h *PongHandler) HandleRequest(_ logrus.FieldLogger, _ *mapleSession.MapleSession, _ *request.RequestReader) {
}
