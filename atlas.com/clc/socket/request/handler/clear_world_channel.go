package handler

import (
	"atlas-clc/mapleSession"
	"atlas-clc/processors"
	"github.com/jtumidanski/atlas-socket/request"
	"log"
)

const OpCodeClearWorldChannel uint16 = 0x0C

type ClearWorldChannelHandler struct {
}

func (h *ClearWorldChannelHandler) IsValid(l *log.Logger, ms *mapleSession.MapleSession) bool {
	v := processors.IsLoggedIn((*ms).AccountId())
	if !v {
		l.Printf("[ERROR] attempting to process a [ClearWorldChannelRequest] when the account %d is not logged in.", (*ms).SessionId())
	}
	return v
}

func (h ClearWorldChannelHandler) HandleRequest(l *log.Logger, ms *mapleSession.MapleSession, _ *request.RequestReader) {
	l.Printf("[INFO] clearing the world and channel for session %d.", (*ms).SessionId())
	(*ms).SetWorldId(0)
	(*ms).SetChannelId(0)
}
