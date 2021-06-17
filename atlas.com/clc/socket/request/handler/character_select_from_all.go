package handler

import (
	"atlas-clc/account"
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

type CharacterSelectFromAllHandler struct {
}

func (h *CharacterSelectFromAllHandler) IsValid(l logrus.FieldLogger, ms *session.MapleSession) bool {
	v := account.IsLoggedIn((*ms).AccountId())
	if !v {
		l.Errorf("Attempting to process a [CharacterSelectFromAllRequest] when the account %d is not logged in.", (*ms).SessionId())
	}
	return v
}

func (h *CharacterSelectFromAllHandler) HandleRequest(l logrus.FieldLogger, ms *session.MapleSession, r *request.RequestReader) {
	p := ReadCharacterSelectFromAll(r)

	c, err := character.GetCharacterById(uint32(p.CharacterId()))
	if err != nil {
		l.WithError(err).Errorf("Unable to retrieve selected character by id")
		return
	}
	if c.Attributes().WorldId() != byte(p.WorldId()) {
		l.Errorf("Client supplied world not matching that of the selected character")
		return
	}
	(*ms).SetWorldId(c.Attributes().WorldId())

	w, err := world.GetWorld((*ms).WorldId())
	if err != nil {
		l.WithError(err).Errorf("Unable to retrieve world logged into by session")
		return
	}
	if w.CapacityStatus() == world.StatusFull {
		l.Infof("World being logged into is full")
		//TODO disconnect
		return
	}

	cs, err := channel.GetChannelsForWorld((*ms).WorldId())
	// initialize global pseudo random generator
	rand.Seed(time.Now().Unix())
	ch := cs[rand.Intn(len(cs))]
	(*ms).SetChannelId(ch.ChannelId())

	(*ms).Announce(writer.WriteServerIp(ch.IpAddress(), ch.Port(), c.Attributes().Id()))
}
