package handlers

import (
	"atlas-clc/models"
	"atlas-clc/packets/inputs"
	"atlas-clc/packets/outputs/writers"
	"atlas-clc/processors"
	"atlas-clc/sessions"
	"log"
)

type ServerListRequestHandler struct {
}

func (h *ServerListRequestHandler) IsValid(l *log.Logger, s *sessions.Session) bool {
	return processors.IsLoggedIn(l, s.AccountId())
}

func (h *ServerListRequestHandler) Handle(l *log.Logger, s *sessions.Session, _ *inputs.Reader) {
	ws, err := processors.GetWorlds(l)
	if err != nil {
		l.Println("[ERROR] retrieving worlds")
		return
	}

	cls, err := processors.GetChannelLoadByWorld(l)
	if err != nil {
		l.Println("[ERROR] retrieving channel load")
		return
	}

	nws := h.combine(l, ws, cls)
	h.respondToSession(s, nws)
}

func (h *ServerListRequestHandler) combine(l *log.Logger, ws []models.World, cls map[int][]models.ChannelLoad) []models.World {
	var nws = make([]models.World, 0)

	for _, x := range ws {
		if val, ok := cls[int(x.Id())]; ok {
			nws = append(nws, *x.SetChannelLoad(val))
		} else {
			l.Println("[ERROR] processing world without a channel load")
			nws = append(nws, x)
		}
	}
	return nws
}

func (h *ServerListRequestHandler) respondToSession(s *sessions.Session, ws []models.World) {
	h.announceServerList(ws, s)
	h.announceLastWorld(s)
	h.announceRecommendedWorlds(ws, s)
}

func (h *ServerListRequestHandler) announceRecommendedWorlds(ws []models.World, s *sessions.Session) {
	var rs = make([]models.WorldRecommendation, 0)
	for _, x := range ws {
		if x.Recommended() {
			rs = append(rs, *x.Recommendation())
		}
	}
	s.Announce(writers.WriteRecommendedWorlds(rs))
}

func (h *ServerListRequestHandler) announceLastWorld(s *sessions.Session) {
	s.Announce(writers.WriteSelectWorld(0))
}

func (h *ServerListRequestHandler) announceServerList(ws []models.World, s *sessions.Session) {
	for _, x := range ws {
		s.Announce(writers.WriteServerListEntry(x.Id(), x.Name(), x.Flag(), x.EventMessage(), x.ChannelLoad()))
	}
	s.Announce(writers.WriteServerListEnd())
}
