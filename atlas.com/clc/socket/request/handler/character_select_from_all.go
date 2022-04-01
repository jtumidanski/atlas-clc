package handler

import (
	"atlas-clc/channel"
	"atlas-clc/character"
	"atlas-clc/session"
	"atlas-clc/socket/response/writer"
	"atlas-clc/world"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const OpCodeCharacterSelectFromAll uint16 = 0x0E

type CharacterSelectFromAllRequest struct {
	characterId int32
	worldId     int32
	macs        string
	hwid        string
}

func (s CharacterSelectFromAllRequest) CharacterId() int32 {
	return s.characterId
}

func (s CharacterSelectFromAllRequest) WorldId() int32 {
	return s.worldId
}

func ReadCharacterSelectFromAll(reader *request.RequestReader) *CharacterSelectFromAllRequest {
	cid := reader.ReadInt32()
	wid := reader.ReadInt32()
	macs := reader.ReadAsciiString()
	hwid := reader.ReadAsciiString()
	return &CharacterSelectFromAllRequest{cid, wid, macs, hwid}
}

func HandleCharacterSelectFromAllRequest(l logrus.FieldLogger, span opentracing.Span) func(s session.Model, r *request.RequestReader) {
	return func(s session.Model, r *request.RequestReader) {
		p := ReadCharacterSelectFromAll(r)

		c, err := character.GetById(l, span)(uint32(p.CharacterId()))
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve selected character by id")
			return
		}
		if c.Properties().WorldId() != byte(p.WorldId()) {
			l.Errorf("Client supplied world not matching that of the selected character")
			return
		}
		s = session.SetWorldId(c.Properties().WorldId())(s.SessionId())

		w, err := world.GetById(l, span)(s.WorldId())
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve world logged into by session")
			return
		}
		if w.CapacityStatus() == world.StatusFull {
			l.Infof("World being logged into is full")
			//TODO disconnect
			return
		}

		ch, err := channel.GetRandomChannelForWorld(l, span)(s.WorldId())
		if err != nil {
			l.WithError(err).Errorf("Unable to select a random channel for the world %d.", s.WorldId())
			//TODO disconnect
			return
		}
		s = session.SetChannelId(ch.ChannelId())(s.SessionId())

		err = session.Announce(writer.WriteServerIp(l)(ch.IpAddress(), ch.Port(), c.Properties().Id()))(s)
		if err != nil {
			l.WithError(err).Errorf("Unable to send channel server connection information")
		}
	}
}
