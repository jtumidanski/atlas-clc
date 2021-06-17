package handler

import (
	"atlas-clc/account"
	"atlas-clc/session"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/sirupsen/logrus"
)

const OpCodeClearWorldChannel uint16 = 0x0C

type ClearWorldChannelHandler struct {
}

func (h *ClearWorldChannelHandler) IsValid(l logrus.FieldLogger, ms *session.MapleSession) bool {
	v := account.IsLoggedIn((*ms).AccountId())
	if !v {
		l.Errorf("Attempting to process a [ClearWorldChannelRequest] when the account %d is not logged in.", (*ms).SessionId())
	}
	return v
}

func (h ClearWorldChannelHandler) HandleRequest(l logrus.FieldLogger, ms *session.MapleSession, _ *request.RequestReader) {
	l.Infof("Clearing the world and channel for session %d.", (*ms).SessionId())
	(*ms).SetWorldId(0)
	(*ms).SetChannelId(0)
}
