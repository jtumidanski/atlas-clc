package handler

import (
	"atlas-clc/session"
	"github.com/jtumidanski/atlas-socket/request"
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

type ClientStartErrorHandler struct {
}

func (c *ClientStartErrorHandler) IsValid(_ logrus.FieldLogger, _ *session.MapleSession) bool {
	return true
}

func (c *ClientStartErrorHandler) HandleRequest(l logrus.FieldLogger, ms *session.MapleSession, r *request.RequestReader) {
	p := ReadClientStartErrorRequest(r)
	l.Errorf("Client start error for %d. Received %s", (*ms).SessionId(), p.Error())
}
