package handler

import (
	status "atlas-clc/domain"
	"atlas-clc/mapleSession"
	"atlas-clc/processors"
	"atlas-clc/socket/response/writer"
	"github.com/jtumidanski/atlas-socket/request"
	"log"
)

const OpCodeCharacterListWorld uint16 = 0x05

type CharacterListWorldRequest struct {
	worldId   byte
	channelId byte
}

func (r CharacterListWorldRequest) WorldId() byte {
	return r.worldId
}

func (r CharacterListWorldRequest) ChannelId() byte {
	return r.channelId
}

func ReadCharacterListWorldRequest(reader *request.RequestReader) *CharacterListWorldRequest {
	reader.ReadByte()
	return &CharacterListWorldRequest{
		worldId:   reader.ReadByte(),
		channelId: reader.ReadByte() + 1,
	}
}

type CharacterListWorldHandler struct {
}

func (h *CharacterListWorldHandler) IsValid(l *log.Logger, ms *mapleSession.MapleSession) bool {
	v := processors.IsLoggedIn((*ms).AccountId())
	if !v {
		l.Printf("[ERROR] attempting to process a [CharacterListWorldRequest] when the account %d is not logged in.", (*ms).SessionId())
	}
	return v
}

func (h *CharacterListWorldHandler) HandleRequest(l *log.Logger, ms *mapleSession.MapleSession, r *request.RequestReader) {
	p := ReadCharacterListWorldRequest(r)

	w, err := processors.GetWorld(p.WorldId())
	if err != nil {
		l.Println("[ERROR] received a character list request for a world we do not have")
		return
	}

	if w.CapacityStatus() == status.Full {
		(*ms).Announce(writer.WriteWorldCapacityStatus(status.Full))
		return
	}

	(*ms).SetWorldId(p.WorldId())
	(*ms).SetChannelId(p.ChannelId())

	a, err := processors.GetAccountById((*ms).AccountId())
	if err != nil {
		l.Println("[ERROR] cannot retrieve account")
		return
	}

	cs, err := processors.GetCharactersForWorld((*ms).AccountId(), p.WorldId())
	if err != nil {
		l.Println("[ERROR] cannot retrieve account characters")
		return
	}

	(*ms).Announce(writer.WriteCharacterList(cs, p.WorldId(), 0, true, a.PIC(), int16(1), a.CharacterSlots()))
}
