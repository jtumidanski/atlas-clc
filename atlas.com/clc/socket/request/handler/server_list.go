package handler

import (
	"atlas-clc/domain"
	"atlas-clc/processors"
	"atlas-clc/sessions"
	"atlas-clc/socket/request"
	"atlas-clc/socket/response/writer"
	"log"
)

const OpCodeServerRequest uint16 = 0x0B
const OpCodeServerListReRequest uint16 = 0x04

type ServerListHandler struct {
}

func (h *ServerListHandler) IsValid(l *log.Logger, s *sessions.Session) bool {
	v := processors.IsLoggedIn(s.AccountId())
	if !v {
		l.Printf("[ERROR] attempting to process a [ServerListRequest] when the account %d is not logged in.", s.SessionId())
	}
	return v
}

func (h *ServerListHandler) HandleRequest(l *log.Logger, s *sessions.Session, _ *request.RequestReader) {
	ws, err := processors.GetWorlds()
	if err != nil {
		l.Println("[ERROR] retrieving worlds")
		return
	}

	cls, err := processors.GetChannelLoadByWorld()
	if err != nil {
		l.Println("[ERROR] retrieving channel load")
		return
	}

	nws := h.combine(l, ws, cls)
	h.respondToSession(s, nws)
}

func (h *ServerListHandler) combine(l *log.Logger, ws []domain.World, cls map[int][]domain.ChannelLoad) []domain.World {
	var nws = make([]domain.World, 0)

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

func (h *ServerListHandler) respondToSession(s *sessions.Session, ws []domain.World) {
	h.announceServerList(ws, s)
	h.announceLastWorld(s)
	h.announceRecommendedWorlds(ws, s)
}

func (h *ServerListHandler) announceRecommendedWorlds(ws []domain.World, s *sessions.Session) {
	var rs = make([]domain.WorldRecommendation, 0)
	for _, x := range ws {
		if x.Recommended() {
			rs = append(rs, *x.Recommendation())
		}
	}
	s.Announce(writer.WriteRecommendedWorlds(rs))
}

func (h *ServerListHandler) announceLastWorld(s *sessions.Session) {
	s.Announce(writer.WriteSelectWorld(0))
}

func (h *ServerListHandler) announceServerList(ws []domain.World, s *sessions.Session) {
	for _, x := range ws {
		s.Announce(writer.WriteServerListEntry(x.Id(), x.Name(), x.Flag(), x.EventMessage(), x.ChannelLoad()))
	}
	s.Announce(writer.WriteServerListEnd())
}
