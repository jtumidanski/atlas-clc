package handler

import (
	"atlas-clc/mapleSession"
	"github.com/jtumidanski/atlas-socket/request"
	"log"
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

func (c *ClientStartErrorHandler) IsValid(_ *log.Logger, _ *mapleSession.MapleSession) bool {
	return true
}

func (c *ClientStartErrorHandler) HandleRequest(l *log.Logger, ms *mapleSession.MapleSession, r *request.RequestReader) {
	p := ReadClientStartErrorRequest(r)
	l.Printf("[ERROR] client start error for %d. Received %s", (*ms).SessionId(), p.Error())
}
