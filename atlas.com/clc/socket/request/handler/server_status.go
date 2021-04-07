package handler

import (
	"atlas-clc/mapleSession"
	"atlas-clc/processors"
	"atlas-clc/socket/response/writer"
	"github.com/jtumidanski/atlas-socket/request"
	"log"
)

const OpCodeServerStatus uint16 = 0x06

type ServerStatusRequest struct {
	worldId byte
}

func (s ServerStatusRequest) WorldId() byte {
	return s.worldId
}

func ReadServerStatusRequest(reader *request.RequestReader) *ServerStatusRequest {
	wid := byte(reader.ReadUint16())
	return &ServerStatusRequest{wid}
}

type ServerStatusHandler struct {
}

func (h *ServerStatusHandler) IsValid(l *log.Logger, ms *mapleSession.MapleSession) bool {
	v := processors.IsLoggedIn((*ms).AccountId())
	if !v {
		l.Printf("[ERROR] attempting to process a [ServerStatusRequest] when the account %d is not logged in.", (*ms).SessionId())
	}
	return v
}

func (h *ServerStatusHandler) HandleRequest(_ *log.Logger, ms *mapleSession.MapleSession, r *request.RequestReader) {
	p := ReadServerStatusRequest(r)

	cs := processors.GetWorldCapacityStatus(p.WorldId())
	(*ms).Announce(writer.WriteWorldCapacityStatus(cs))
}