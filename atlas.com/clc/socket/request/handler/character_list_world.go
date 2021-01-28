package handler

import (
	status "atlas-clc/domain"
	"atlas-clc/processors"
	"atlas-clc/sessions"
	"atlas-clc/socket/request"
	"atlas-clc/socket/response/writer"
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

func (h *CharacterListWorldHandler) IsValid(l *log.Logger, s *sessions.Session) bool {
	v := processors.IsLoggedIn(s.AccountId())
	if !v {
		l.Printf("[ERROR] attempting to process a [CharacterListWorldRequest] when the account %d is not logged in.", s.SessionId())
	}
	return v
}

func (h *CharacterListWorldHandler) HandleRequest(l *log.Logger, s *sessions.Session, r *request.RequestReader) {
	p := ReadCharacterListWorldRequest(r)

	w, err := processors.GetWorld(p.WorldId())
	if err != nil {
		l.Println("[ERROR] received a character list request for a world we do not have")
		return
	}

	if w.CapacityStatus() == status.Full {
		s.Announce(writer.WriteWorldCapacityStatus(status.Full))
		return
	}

	s.SetWorldId(p.WorldId())
	s.SetChannelId(p.ChannelId())

	a, err := processors.GetAccountById(s.AccountId())
	if err != nil {
		l.Println("[ERROR] cannot retrieve account")
		return
	}

	cs, err := processors.GetCharactersForWorld(s.AccountId(), p.WorldId())
	if err != nil {
		l.Println("[ERROR] cannot retrieve account characters")
		return
	}

	s.Announce(writer.WriteCharacterList(cs, p.WorldId(), 0, true, a.PIC(), int16(1), a.CharacterSlots()))
}
