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

func HandleCharacterListWorldRequest(l logrus.FieldLogger, ms *session.MapleSession, r *request.RequestReader) {
	p := ReadCharacterListWorldRequest(r)

	w, err := world.GetById(p.WorldId())
	if err != nil {
		l.WithError(err).Errorf("Received a character list request for a world we do not have")
		return
	}

	if w.CapacityStatus() == world.StatusFull {
		err = (*ms).Announce(writer.WriteWorldCapacityStatus(l)(world.StatusFull))
		if err != nil {
			l.WithError(err).Errorf("Unable to show that world %d is full", w.Id())
		}
		return
	}

	(*ms).SetWorldId(p.WorldId())
	(*ms).SetChannelId(p.ChannelId())

	a, err := account.GetById((*ms).AccountId())
	if err != nil {
		l.WithError(err).Errorf("Cannot retrieve account")
		return
	}

	cs, err := character.GetForWorld((*ms).AccountId(), p.WorldId())
	if err != nil {
		l.WithError(err).Errorf("Cannot retrieve account characters")
		return
	}

	err = (*ms).Announce(writer.WriteCharacterList(l)(cs, p.WorldId(), 0, true, a.PIC(), int16(1), a.CharacterSlots()))
	if err != nil {
		l.WithError(err).Errorf("Unable to show character list")
	}
}
