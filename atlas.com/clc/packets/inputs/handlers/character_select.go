package handlers

import (
	models2 "atlas-clc/models"
	"atlas-clc/packets/inputs"
	"atlas-clc/packets/inputs/models"
	"atlas-clc/packets/inputs/readers"
	"atlas-clc/packets/outputs/writers"
	"atlas-clc/processors"
	"atlas-clc/registries"
	"atlas-clc/sessions"
	"log"
)

type CharacterSelectHandler struct {
}

func (h *CharacterSelectHandler) Handle(l *log.Logger, sessionId int, r *inputs.Reader) {
	s := registries.GetSessionRegistry().GetSession(sessionId)
	if s != nil {
		p := readers.ReadCharacterSelected(r)
		if p != nil {
			h.handle(l, s, p)
		}
	}
}

func (h *CharacterSelectHandler) handle(l *log.Logger, s *sessions.Session, p *models.CharacterSelected) {

	c, err := processors.GetCharacterById(l, uint32(p.CharacterId()))
	if err != nil {
		l.Println("[ERROR] unable to retrieve selected character by id")
		return
	}

	w, err := processors.GetWorld(l, s.WorldId())
	if err != nil {
		l.Println("[ERROR] unable to retrieve world logged into by session")
		return
	}
	if w.CapacityStatus() == models2.Full {
		l.Println("[INFO] world being logged into is full")
		//TODO disconnect
		return
	}

	ch, err := processors.GetChannelForWorld(l, s.WorldId(), s.ChannelId())
	if err != nil {
		l.Println("[ERROR] unable to retrieve channel in world")
		return
	}

	s.Announce(writers.WriteServerIp(ch.IpAddress(), ch.Port(), c.Attributes().Id()))
}
