package request

import (
	"atlas-clc/session"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/sirupsen/logrus"
)

type HandlerSupplier struct {
	l logrus.FieldLogger
}

type MapleHandler interface {
	IsValid(l logrus.FieldLogger, s *session.MapleSession) bool

	HandleRequest(l logrus.FieldLogger, s *session.MapleSession, r *request.RequestReader)
}

func AdaptHandler(l logrus.FieldLogger, h MapleHandler) func(int, request.RequestReader) {
	return func(sessionId int, r request.RequestReader) {
		s := session.GetRegistry().Get(sessionId)
		if s != nil {
			if h.IsValid(l, &s) {
				h.HandleRequest(l, &s, &r)
				s.UpdateLastRequest()
			}
		} else {
			l.Errorf("Unable to locate session %d", sessionId)
		}
	}
}
