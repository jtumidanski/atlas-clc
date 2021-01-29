package sessions

import (
	"atlas-clc/registries"
	session2 "github.com/jtumidanski/atlas-socket/session"
	"io"
	"log"
	"net"
)

type SessionService interface {
	session2.Service
}

type sessionService struct {
	l *log.Logger
	r *registries.SessionRegistry
}

func NewSessionService(l *log.Logger) SessionService {
	return &sessionService{l, registries.GetSessionRegistry()}
}

func NewReaderService() io.Reader {
	var rw io.ReadWriter
	c := rw.(io.Reader)
	return c
}

func (s *sessionService) Create(l *log.Logger, sessionId int, conn net.Conn) (session2.Session, error) {
	session := NewSession(sessionId, conn, l)
	s.r.Add(&session)
	return session, nil
}

func (s *sessionService) Get(sessionId int) session2.Session {
	return s.r.Get(sessionId)
}

func (s *sessionService) GetAll() []session2.Session {
	ss := s.r.GetAll()
	b := make([]session2.Session, len(ss))
	for i, v := range ss {
		b[i] = v.(session2.Session)
	}
	return b
}

func (s *sessionService) Destroy(sessionId int) {
	s.r.Remove(sessionId)
}
