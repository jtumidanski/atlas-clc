package services

import (
	"atlas-clc/character"
	session2 "atlas-clc/session"
	"github.com/jtumidanski/atlas-socket/session"
	"github.com/sirupsen/logrus"
	"net"
)

type Service interface {
	session.Service
}

type mapleSessionService struct {
	l logrus.FieldLogger
	r *session2.SessionRegistry
}

func NewMapleSessionService(l logrus.FieldLogger) Service {
	return &mapleSessionService{l, session2.GetSessionRegistry()}
}

func (s *mapleSessionService) Create(sessionId int, conn net.Conn) (session.Session, error) {
	ses := session2.NewSession(sessionId, conn)
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
	ses := s.Get(sessionId).(session2.MapleSession)

	s.r.Remove(sessionId)

	character.Logout(s.l)(ses.WorldId(), ses.ChannelId(), ses.AccountId(), 0)
}
