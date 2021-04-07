package handler

import (
	"atlas-clc/kafka/producers"
	"atlas-clc/mapleSession"
	"atlas-clc/processors"
	"context"
	"github.com/jtumidanski/atlas-socket/request"
	"log"
)

const OpCodeDisconnect uint16 = 0x0C

type DisconnectHandler struct {
}

func (h *DisconnectHandler) IsValid(l *log.Logger, ms *mapleSession.MapleSession) bool {
	v := processors.IsLoggedIn((*ms).AccountId())
	if !v {
		l.Printf("[ERROR] attempting to process a [Disconnect] when the account %d is not logged in.", (*ms).SessionId())
	}
	return v
}

func (h DisconnectHandler) HandleRequest(l *log.Logger, ms *mapleSession.MapleSession, _ *request.RequestReader) {
	producers.CharacterStatus(l, context.Background()).Logout((*ms).WorldId(), (*ms).ChannelId(), (*ms).AccountId(), 0)
}
