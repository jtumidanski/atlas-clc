package handler

import (
	"atlas-clc/session"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/sirupsen/logrus"
)

const OpCodePong uint16 = 0x18

type PongHandler struct {
}

func (h *PongHandler) IsValid(_ logrus.FieldLogger, _ *session.MapleSession) bool {
	return true
}

func (h *PongHandler) HandleRequest(_ logrus.FieldLogger, _ *session.MapleSession, _ *request.RequestReader) {
}
