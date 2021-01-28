package sessions

import (
	"atlas-clc/registries"
	"atlas-clc/socket"
	"log"
	"net"
)

type SessionService interface {
	socket.SocketSessionCreator
}

type sessionService struct {
	l *log.Logger
	r *registries.SessionRegistry
}

func NewSessionService(l *log.Logger) SessionService {
	return sessionService{l, registries.GetSessionRegistry()}
}

func (s *sessionService) Create(sessionId int, conn *net.Conn, l *log.Logger) (*Session, error) {
	return NewSession(sessionId, conn, l), nil
}

func (s *sessionService) Add(session session) {
	s.r.Add(session)
}

func (s *sessionService) Get(sessionId int) Session {
	return s.r.Get(sessionId)
}

func (s *sessionService) GetAll() []Session {
	return s.r.GetAll()
}

func (s *sessionService) Destroy(sessionId int) {
	s.r.Remove(sessionId)
}
