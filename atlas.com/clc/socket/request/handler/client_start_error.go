package handler

import (
	"atlas-clc/session"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const OpCodeClientStartError uint16 = 0x19

type ClientStartErrorRequest struct {
	error string
}

func (c *ClientStartErrorRequest) Error() string {
	return c.error
}

func ReadClientStartErrorRequest(reader *request.RequestReader) *ClientStartErrorRequest {
	message := reader.ReadAsciiString()
	return &ClientStartErrorRequest{message}
}

func HandleClientStartErrorRequest(l logrus.FieldLogger, _ opentracing.Span) func(s session.Model, r *request.RequestReader) {
	return func(s session.Model, r *request.RequestReader) {
		p := ReadClientStartErrorRequest(r)
		l.Errorf("Client start error for %d. Received %s", s.SessionId(), p.Error())
	}
}
