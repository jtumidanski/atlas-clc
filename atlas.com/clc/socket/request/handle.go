package request

import (
	"atlas-clc/account"
	"atlas-clc/session"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/sirupsen/logrus"
)

type MessageValidator func(l logrus.FieldLogger, s *session.MapleSession) bool

func NoOpValidator(_ logrus.FieldLogger, _ *session.MapleSession) bool {
	return true
}

func LoggedInValidator(l logrus.FieldLogger, s *session.MapleSession) bool {
	v := account.IsLoggedIn((*s).AccountId())
	if !v {
		l.Errorf("Attempting to process a request when the account %d is not logged in.", (*s).SessionId())
	}
	return v
}

type MessageHandler func(l logrus.FieldLogger, s *session.MapleSession, r *request.RequestReader)

func NoOpHandler(_ logrus.FieldLogger, _ *session.MapleSession, _ *request.RequestReader) {
}

func AdaptHandler(l logrus.FieldLogger, v MessageValidator, h MessageHandler) func(int, request.RequestReader) {
	return func(sessionId int, r request.RequestReader) {
		s := session.GetRegistry().Get(sessionId)
		if s != nil {
			if v(l, &s) {
				h(l, &s, &r)
				s.UpdateLastRequest()
			}
		} else {
			l.Errorf("Unable to locate session %d", sessionId)
		}
	}
}
