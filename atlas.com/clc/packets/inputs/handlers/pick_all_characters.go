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
	"math/rand"
	"time"
)

type PickAllCharactersHandler struct {
}

func (h *PickAllCharactersHandler) Handle(l *log.Logger, sessionId int, r *inputs.Reader) {
	s := registries.GetSessionRegistry().GetSession(sessionId)
	if s != nil {
		p := readers.ReadViewAllCharactersSelected(r)
		if p != nil {
			h.handle(l, s, p)
		}
	}
}

func (h *PickAllCharactersHandler) handle(l *log.Logger, s *sessions.Session, p *models.ViewAllCharactersSelected) {
	c, err := processors.GetCharacterById(l, uint32(p.CharacterId()))
	if err != nil {
		l.Println("[ERROR] unable to retrieve selected character by id")
		return
	}
	if c.Attributes().WorldId() != byte(p.WorldId()) {
		l.Println("[ERROR] client supplied world not matching that of the selected character")
		return
	}
	s.SetWorldId(c.Attributes().WorldId())

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

	cs, err := processors.GetChannelsForWorld(l, s.WorldId())
	// initialize global pseudo random generator
	rand.Seed(time.Now().Unix())
	ch := cs[rand.Intn(len(cs))]
	s.SetChannelId(ch.ChannelId())

	s.Announce(writers.WriteServerIp(ch.IpAddress(), ch.Port(), c.Attributes().Id()))
}
