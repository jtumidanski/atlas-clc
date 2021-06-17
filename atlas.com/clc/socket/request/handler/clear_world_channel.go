package handler

import (
	"atlas-clc/session"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/sirupsen/logrus"
)

const OpCodeClearWorldChannel uint16 = 0x0C

func HandleClearWorldChannelRequest(l logrus.FieldLogger, ms *session.Model, _ *request.RequestReader) {
	l.Infof("Clearing the world and channel for session %d.", ms.SessionId())
	ms.SetWorldId(0)
	ms.SetChannelId(0)
}
