package handler

import (
	"atlas-clc/account"
	"atlas-clc/channel"
	"atlas-clc/session"
	"atlas-clc/socket/response/writer"
	"atlas-clc/world"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/sirupsen/logrus"
)

const OpCodeServerRequest uint16 = 0x0B
const OpCodeServerListReRequest uint16 = 0x04

type ServerListHandler struct {
}

func (h *ServerListHandler) IsValid(l logrus.FieldLogger, ms *session.MapleSession) bool {
	v := account.IsLoggedIn((*ms).AccountId())
	if !v {
		l.Errorf("Attempting to process a [ServerListRequest] when the account %d is not logged in.", (*ms).SessionId())
	}
	return v
}

func (h *ServerListHandler) HandleRequest(l logrus.FieldLogger, ms *session.MapleSession, _ *request.RequestReader) {
	ws, err := world.GetAll()
	if err != nil {
		l.WithError(err).Errorf("Retrieving worlds")
		return
	}

	cls, err := channel.GetChannelLoadByWorld()
	if err != nil {
		l.WithError(err).Errorf("Retrieving channel load")
		return
	}

	nws := h.combine(l, ws, cls)
	h.respondToSession(l, ms, nws)
}

func (h *ServerListHandler) combine(l logrus.FieldLogger, ws []world.Model, cls map[int][]channel.Load) []world.Model {
	var nws = make([]world.Model, 0)

	for _, x := range ws {
		if val, ok := cls[int(x.Id())]; ok {
			nws = append(nws, x.SetChannelLoad(val))
		} else {
			l.Errorf("Processing world without a channel load")
			nws = append(nws, x)
		}
	}
	return nws
}

func (h *ServerListHandler) respondToSession(l logrus.FieldLogger, ms *session.MapleSession, ws []world.Model) {
	h.announceServerList(l, ws, ms)
	h.announceLastWorld(l, ms)
	h.announceRecommendedWorlds(l, ws, ms)
}

func (h *ServerListHandler) announceRecommendedWorlds(l logrus.FieldLogger, ws []world.Model, ms *session.MapleSession) {
	var rs = make([]world.Recommendation, 0)
	for _, x := range ws {
		if x.Recommended() {
			rs = append(rs, x.Recommendation())
		}
	}
	err := (*ms).Announce(writer.WriteRecommendedWorlds(rs))
	if err != nil {
		l.WithError(err).Errorf("Unable to issue recommended worlds")
	}
}

func (h *ServerListHandler) announceLastWorld(l logrus.FieldLogger, ms *session.MapleSession) {
	err := (*ms).Announce(writer.WriteSelectWorld(0))
	if err != nil {
		l.WithError(err).Errorf("Unable to identify the last world a account was logged into")
	}
}

func (h *ServerListHandler) announceServerList(l logrus.FieldLogger, ws []world.Model, ms *session.MapleSession) {
	for _, x := range ws {
		err := (*ms).Announce(writer.WriteServerListEntry(x.Id(), x.Name(), x.Flag(), x.EventMessage(), x.ChannelLoad()))
		if err != nil {
			l.WithError(err).Errorf("Unable to write server list entry for %d", x.Id())
		}
	}
	err := (*ms).Announce(writer.WriteServerListEnd())
	if err != nil {
		l.WithError(err).Errorf("Unable to complete writing the server list")
	}
}
