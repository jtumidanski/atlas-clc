package handler

import (
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

func HandleServerStatusRequest(l logrus.FieldLogger, ms *session.Model, r *request.RequestReader) {
	p := ReadServerStatusRequest(r)

	cs := world.GetCapacityStatus(l)(p.WorldId())
	err := ms.Announce(writer.WriteWorldCapacityStatus(l)(cs))
	if err != nil {
		l.WithError(err).Errorf("Unable to issue world capacity status information")
	}
}
