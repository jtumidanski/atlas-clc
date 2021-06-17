package session

import (
	"atlas-clc/character"
	"github.com/jtumidanski/atlas-socket/session"
	"github.com/sirupsen/logrus"
	"net"
)

type Service interface {
	session.Service
}

type mapleSessionService struct {
	l logrus.FieldLogger
	r *Registry
}

func NewMapleSessionService(l logrus.FieldLogger) Service {
	return &mapleSessionService{l, GetRegistry()}
}

func (s *mapleSessionService) Create(sessionId int, conn net.Conn) (session.Session, error) {
	ses := NewSession(sessionId, conn)
	s.r.Add(&ses)
	ses.WriteHello()
	return ses, nil
}

func (s *mapleSessionService) Get(sessionId int) session.Session {
	return s.r.Get(sessionId)
}

func (s *mapleSessionService) GetAll() []session.Session {
	ss := s.r.GetAll()
	b := make([]session.Session, len(ss))
	for i, v := range ss {
		b[i] = v.(session.Session)
	}
	return b
}

func (s *mapleSessionService) Destroy(sessionId int) {
	ses := s.Get(sessionId).(MapleSession)

	s.r.Remove(sessionId)

	character.Logout(s.l)(ses.WorldId(), ses.ChannelId(), ses.AccountId(), 0)
}
