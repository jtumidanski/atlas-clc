package handler

import (
	"atlas-clc/channel"
	"atlas-clc/character"
	"atlas-clc/session"
	"atlas-clc/socket/response/writer"
	"atlas-clc/world"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
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

func HandleCharacterSelectFromAllRequest(l logrus.FieldLogger, ms *session.Model, r *request.RequestReader) {
	p := ReadCharacterSelectFromAll(r)

	c, err := character.GetById(l)(uint32(p.CharacterId()))
	if err != nil {
		l.WithError(err).Errorf("Unable to retrieve selected character by id")
		return
	}
	if c.Properties().WorldId() != byte(p.WorldId()) {
		l.Errorf("Client supplied world not matching that of the selected character")
		return
	}
	ms.SetWorldId(c.Properties().WorldId())

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

	cs, err := channel.GetAllForWorld(l)(ms.WorldId())
	// initialize global pseudo random generator
	rand.Seed(time.Now().Unix())
	ch := cs[rand.Intn(len(cs))]
	ms.SetChannelId(ch.ChannelId())

	err = ms.Announce(writer.WriteServerIp(l)(ch.IpAddress(), ch.Port(), c.Properties().Id()))
	if err != nil {
		l.WithError(err).Errorf("Unable to send channel server connection information")
	}
}
