package handler

import (
	"atlas-clc/channel"
	"atlas-clc/character"
	"atlas-clc/session"
	"atlas-clc/socket/response/writer"
	"atlas-clc/world"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/sirupsen/logrus"
)

const OpCodeCharacterSelectFromWorld uint16 = 0x13

type CharacterSelectFromWorldRequest struct {
	characterId int32
	macs        string
	hwid        string
}

func (s CharacterSelectFromWorldRequest) CharacterId() int32 {
	return s.characterId
}

func ReadCharacterSelectFromWorldRequest(reader *request.RequestReader) *CharacterSelectFromWorldRequest {
	cid := reader.ReadInt32()
	macs := reader.ReadAsciiString()
	hwid := reader.ReadAsciiString()

	return &CharacterSelectFromWorldRequest{cid, macs, hwid}
}

func HandleCharacterSelectFromWorldRequest(l logrus.FieldLogger, ms *session.Model, r *request.RequestReader) {
	p := ReadCharacterSelectFromWorldRequest(r)

	c, err := character.GetById(l)(uint32(p.CharacterId()))
	if err != nil {
		l.WithError(err).Errorf("Unable to retrieve selected character by id")
		return
	}

	w, err := world.GetById(l)(ms.WorldId())
	if err != nil {
		l.WithError(err).Errorf("Unable to retrieve world logged into by session")
		return
	}
	if w.CapacityStatus() == world.StatusFull {
		l.Infof("World being logged into is full")
		//TODO disconnect
		return
	}

	ch, err := channel.GetForWorldById(l)(ms.WorldId(), ms.ChannelId())
	if err != nil {
		l.WithError(err).Errorf("Unable to retrieve channel in world")
		return
	}

	err = ms.Announce(writer.WriteServerIp(l)(ch.IpAddress(), ch.Port(), c.Properties().Id()))
	if err != nil {
		l.WithError(err).Errorf("Unable to send channel server connection information")
	}
}
