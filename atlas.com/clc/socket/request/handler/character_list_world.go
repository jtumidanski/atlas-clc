package handler

import (
	status "atlas-clc/domain"
	"atlas-clc/mapleSession"
	"atlas-clc/processors"
	"atlas-clc/socket/response/writer"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/sirupsen/logrus"
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

func (h *CharacterListWorldHandler) IsValid(l logrus.FieldLogger, ms *mapleSession.MapleSession) bool {
	v := processors.IsLoggedIn((*ms).AccountId())
	if !v {
		l.Errorf("Attempting to process a [CharacterListWorldRequest] when the account %d is not logged in.", (*ms).SessionId())
	}
	return v
}

func (h *CharacterListWorldHandler) HandleRequest(l logrus.FieldLogger, ms *mapleSession.MapleSession, r *request.RequestReader) {
	p := ReadCharacterListWorldRequest(r)

	w, err := processors.GetWorld(p.WorldId())
	if err != nil {
		l.WithError(err).Errorf("Received a character list request for a world we do not have")
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
		l.WithError(err).Errorf("Cannot retrieve account")
		return
	}

	cs, err := processors.GetCharactersForWorld((*ms).AccountId(), p.WorldId())
	if err != nil {
		l.WithError(err).Errorf("Cannot retrieve account characters")
		return
	}

	(*ms).Announce(writer.WriteCharacterList(cs, p.WorldId(), 0, true, a.PIC(), int16(1), a.CharacterSlots()))
}
