package handler

import (
	"atlas-clc/account"
	"atlas-clc/session"
	"atlas-clc/socket/response/writer"
	"atlas-clc/world"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/sirupsen/logrus"
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

func (h *ServerStatusHandler) IsValid(l logrus.FieldLogger, ms *session.MapleSession) bool {
	v := account.IsLoggedIn((*ms).AccountId())
	if !v {
		l.Errorf("Attempting to process a [ServerStatusRequest] when the account %d is not logged in.", (*ms).SessionId())
	}
	return v
}

func (h *ServerStatusHandler) HandleRequest(l logrus.FieldLogger, ms *session.MapleSession, r *request.RequestReader) {
	p := ReadServerStatusRequest(r)

	cs := world.GetWorldCapacityStatus(p.WorldId())
	err := (*ms).Announce(writer.WriteWorldCapacityStatus(l)(cs))
	if err != nil {
		l.WithError(err).Errorf("Unable to issue world capacity status information")
	}
}
