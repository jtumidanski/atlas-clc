package handler

import (
	models2 "atlas-clc/domain"
	"atlas-clc/mapleSession"
	"atlas-clc/processors"
	"atlas-clc/socket/response/writer"
	"github.com/jtumidanski/atlas-socket/request"
	"log"
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

func (h *CharacterSelectFromAllHandler) IsValid(l *log.Logger, ms *mapleSession.MapleSession) bool {
	v := processors.IsLoggedIn((*ms).AccountId())
	if !v {
		l.Printf("[ERROR] attempting to process a [CharacterSelectFromAllRequest] when the account %d is not logged in.", (*ms).SessionId())
	}
	return v
}

func (h *CharacterSelectFromAllHandler) HandleRequest(l *log.Logger, ms *mapleSession.MapleSession, r *request.RequestReader) {
	p := ReadCharacterSelectFromAll(r)

	c, err := processors.GetCharacterById(uint32(p.CharacterId()))
	if err != nil {
		l.Println("[ERROR] unable to retrieve selected character by id")
		return
	}
	if c.Attributes().WorldId() != byte(p.WorldId()) {
		l.Println("[ERROR] client supplied world not matching that of the selected character")
		return
	}
	(*ms).SetWorldId(c.Attributes().WorldId())

	w, err := processors.GetWorld((*ms).WorldId())
	if err != nil {
		l.Println("[ERROR] unable to retrieve world logged into by session")
		return
	}
	if w.CapacityStatus() == models2.Full {
		l.Println("[INFO] world being logged into is full")
		//TODO disconnect
		return
	}

	cs, err := processors.GetChannelsForWorld((*ms).WorldId())
	// initialize global pseudo random generator
	rand.Seed(time.Now().Unix())
	ch := cs[rand.Intn(len(cs))]
	(*ms).SetChannelId(ch.ChannelId())

	(*ms).Announce(writer.WriteServerIp(ch.IpAddress(), ch.Port(), c.Attributes().Id()))
}
