package services

import (
	"atlas-clc/mapleSession"
	"atlas-clc/registries"
	"github.com/jtumidanski/atlas-socket/session"
	"net"
)

type Service interface {
	session.Service
}

type mapleSessionService struct {
	r *registries.SessionRegistry
}

func NewMapleSessionService() Service {
	return &mapleSessionService{registries.GetSessionRegistry()}
}

func (s *mapleSessionService) Create(sessionId int, conn net.Conn) (session.Session, error) {
	ses := mapleSession.NewSession(sessionId, conn)
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
	s.r.Remove(sessionId)
}
