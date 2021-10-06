package handler

import (
	"atlas-clc/session"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const OpCodeClearWorldChannel uint16 = 0x0C

func HandleClearWorldChannelRequest(l logrus.FieldLogger, _ opentracing.Span) func(s *session.Model, _ *request.RequestReader) {
	return func(s *session.Model, _ *request.RequestReader) {
		l.Infof("Clearing the world and channel for session %d.", s.SessionId())
		s.SetWorldId(0)
		s.SetChannelId(0)
	}
}
