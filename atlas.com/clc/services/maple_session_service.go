package services

import (
	"atlas-clc/mapleSession"
	"atlas-clc/registries"
	"github.com/jtumidanski/atlas-socket/session"
	"log"
	"net"
)

type Service interface {
	session.Service
}

type mapleSessionService struct {
	l *log.Logger
	r *registries.SessionRegistry
}

func NewMapleSessionService(l *log.Logger) Service {
	return &mapleSessionService{l, registries.GetSessionRegistry()}
}

func (s *mapleSessionService) Create(l *log.Logger, sessionId int, conn net.Conn) (session.Session, error) {
	ses := mapleSession.NewSession(sessionId, conn, l)
	s.r.Add(&ses)
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
	s.r.Remove(sessionId)
}
