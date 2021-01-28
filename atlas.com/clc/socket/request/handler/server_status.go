package handler

import (
	"atlas-clc/processors"
	"atlas-clc/sessions"
	"atlas-clc/socket/request"
	"atlas-clc/socket/response/writer"
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

func (h *ServerStatusHandler) IsValid(l *log.Logger, s *sessions.Session) bool {
	v := processors.IsLoggedIn(s.AccountId())
	if !v {
		l.Printf("[ERROR] attempting to process a [ServerStatusRequest] when the account %d is not logged in.", s.SessionId())
	}
	return v
}

func (h *ServerStatusHandler) HandleRequest(_ *log.Logger, s *sessions.Session, r *request.RequestReader) {
	p := ReadServerStatusRequest(r)

	cs := processors.GetWorldCapacityStatus(p.WorldId())
	s.Announce(writer.WriteWorldCapacityStatus(cs))
}
