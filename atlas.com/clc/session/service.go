package session

import (
	"atlas-clc/character"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"net"
)

func Create(l logrus.FieldLogger, r *Registry) func(sessionId uint32, conn net.Conn) {
	return func(sessionId uint32, conn net.Conn) {
		l.Debugf("Creating session %d.", sessionId)
		s := NewSession(sessionId, conn)
		r.Add(s)
		s.WriteHello()
	}
}

func Decrypt(_ logrus.FieldLogger, r *Registry) func(sessionId uint32, input []byte) []byte {
	return func(sessionId uint32, input []byte) []byte {
		s := r.Get(sessionId)
		if s == nil {
			return input
		}
		if s.ReceiveAESOFB() == nil {
			return input
		}
		return s.ReceiveAESOFB().Decrypt(input, true, true)
	}
}

func DestroyAll(l logrus.FieldLogger, span opentracing.Span, r *Registry) {
	for _, s := range r.GetAll() {
		Destroy(l, span, r)(&s)
	}
}

func DestroyByIdWithSpan(l logrus.FieldLogger, r *Registry) func(sessionId uint32) {
	return func(sessionId uint32) {
		span := opentracing.StartSpan("session_destroy")
		defer span.Finish()
		DestroyById(l, span, r)(sessionId)
	}
}

func DestroyById(l logrus.FieldLogger, span opentracing.Span, r *Registry) func(sessionId uint32) {
	return func(sessionId uint32) {
		s := r.Get(sessionId)
		if s == nil {
			return
		}
		Destroy(l, span, r)(s)
	}
}

func Destroy(l logrus.FieldLogger, span opentracing.Span, r *Registry) func(*Model) {
	return func(s *Model) {
		l.Debugf("Destroying session %d.", s.SessionId())
		r.Remove(s.SessionId())
		character.Logout(l, span)(s.WorldId(), s.ChannelId(), s.AccountId(), 0)
	}
}
