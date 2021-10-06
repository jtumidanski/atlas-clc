package handler

import (
	"atlas-clc/account"
	"atlas-clc/character"
	"atlas-clc/session"
	"atlas-clc/socket/response/writer"
	"atlas-clc/world"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/opentracing/opentracing-go"
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

func HandleCharacterListWorldRequest(l logrus.FieldLogger, span opentracing.Span) func(s *session.Model, r *request.RequestReader) {
	return func(s *session.Model, r *request.RequestReader) {
		p := ReadCharacterListWorldRequest(r)

		w, err := world.GetById(l, span)(p.WorldId())
		if err != nil {
			l.WithError(err).Errorf("Received a character list request for a world we do not have")
			return
		}

		if w.CapacityStatus() == world.StatusFull {
			err = s.Announce(writer.WriteWorldCapacityStatus(l)(world.StatusFull))
			if err != nil {
				l.WithError(err).Errorf("Unable to show that world %d is full", w.Id())
			}
			return
		}

		s.SetWorldId(p.WorldId())
		s.SetChannelId(p.ChannelId())

		a, err := account.GetById(l, span)(s.AccountId())
		if err != nil {
			l.WithError(err).Errorf("Cannot retrieve account")
			return
		}

		cs, err := character.GetForWorld(l, span)(s.AccountId(), p.WorldId())
		if err != nil {
			l.WithError(err).Errorf("Cannot retrieve account characters")
			return
		}

		err = s.Announce(writer.WriteCharacterList(l)(cs, p.WorldId(), 0, true, a.PIC(), int16(1), a.CharacterSlots()))
		if err != nil {
			l.WithError(err).Errorf("Unable to show character list")
		}
	}
}
