package handler

import (
	models2 "atlas-clc/domain"
	"atlas-clc/mapleSession"
	"atlas-clc/processors"
	"atlas-clc/socket/response/writer"
	"github.com/jtumidanski/atlas-socket/request"
	"log"
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

type CharacterSelectFromWorldHandler struct {
}

func (h *CharacterSelectFromWorldHandler) IsValid(l *log.Logger, ms *mapleSession.MapleSession) bool {
	v := processors.IsLoggedIn((*ms).AccountId())
	if !v {
		l.Printf("[ERROR] attempting to process a [CharacterSelectFromWorldRequest] when the account %d is not logged in.", (*ms).SessionId())
	}
	return v
}

func (h *CharacterSelectFromWorldHandler) HandleRequest(l *log.Logger, ms *mapleSession.MapleSession, r *request.RequestReader) {
	p := ReadCharacterSelectFromWorldRequest(r)

	c, err := processors.GetCharacterById(uint32(p.CharacterId()))
	if err != nil {
		l.Println("[ERROR] unable to retrieve selected character by id")
		return
	}

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

	ch, err := processors.GetChannelForWorld((*ms).WorldId(), (*ms).ChannelId())
	if err != nil {
		l.Println("[ERROR] unable to retrieve channel in world")
		return
	}

	(*ms).Announce(writer.WriteServerIp(ch.IpAddress(), ch.Port(), c.Attributes().Id()))
}
