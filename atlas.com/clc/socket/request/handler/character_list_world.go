package handler

import (
	"atlas-clc/account"
	"atlas-clc/character"
	"atlas-clc/session"
	"atlas-clc/socket/response/writer"
	"atlas-clc/world"
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

func (h *CharacterListWorldHandler) IsValid(l logrus.FieldLogger, ms *session.MapleSession) bool {
	v := account.IsLoggedIn((*ms).AccountId())
	if !v {
		l.Errorf("Attempting to process a [CharacterListWorldRequest] when the account %d is not logged in.", (*ms).SessionId())
	}
	return v
}

func (h *CharacterListWorldHandler) HandleRequest(l logrus.FieldLogger, ms *session.MapleSession, r *request.RequestReader) {
	p := ReadCharacterListWorldRequest(r)

	w, err := world.GetWorld(p.WorldId())
	if err != nil {
		l.WithError(err).Errorf("Received a character list request for a world we do not have")
		return
	}

	if w.CapacityStatus() == world.StatusFull {
		(*ms).Announce(writer.WriteWorldCapacityStatus(world.StatusFull))
		return
	}

	(*ms).SetWorldId(p.WorldId())
	(*ms).SetChannelId(p.ChannelId())

	a, err := account.GetAccountById((*ms).AccountId())
	if err != nil {
		l.WithError(err).Errorf("Cannot retrieve account")
		return
	}

	cs, err := character.GetCharactersForWorld((*ms).AccountId(), p.WorldId())
	if err != nil {
		l.WithError(err).Errorf("Cannot retrieve account characters")
		return
	}

	(*ms).Announce(writer.WriteCharacterList(cs, p.WorldId(), 0, true, a.PIC(), int16(1), a.CharacterSlots()))
}
