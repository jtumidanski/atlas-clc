package handler

import (
	"atlas-clc/channel"
	"atlas-clc/session"
	"atlas-clc/socket/response/writer"
	"atlas-clc/world"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/sirupsen/logrus"
)

const OpCodeServerRequest uint16 = 0x0B
const OpCodeServerListReRequest uint16 = 0x04

func HandleServerListRequest(l logrus.FieldLogger, ms *session.Model, _ *request.RequestReader) {
	ws, err := world.GetAll(l)
	if err != nil {
		l.WithError(err).Errorf("Retrieving worlds")
		return
	}

	cls, err := channel.GetChannelLoadByWorld(l)
	if err != nil {
		l.WithError(err).Errorf("Retrieving channel load")
		return
	}

	nws := combine(l, ws, cls)
	respondToSession(l, ms, nws)
}

func combine(l logrus.FieldLogger, ws []world.Model, cls map[int][]channel.Load) []world.Model {
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

func respondToSession(l logrus.FieldLogger, ms *session.Model, ws []world.Model) {
	announceServerList(l, ws, ms)
	announceLastWorld(l, ms)
	announceRecommendedWorlds(l, ws, ms)
}

func announceRecommendedWorlds(l logrus.FieldLogger, ws []world.Model, ms *session.Model) {
	var rs = make([]world.Recommendation, 0)
	for _, x := range ws {
		if x.Recommended() {
			rs = append(rs, x.Recommendation())
		}
	}
	err := ms.Announce(writer.WriteRecommendedWorlds(l)(rs))
	if err != nil {
		l.WithError(err).Errorf("Unable to issue recommended worlds")
	}
}

func announceLastWorld(l logrus.FieldLogger, ms *session.Model) {
	err := ms.Announce(writer.WriteSelectWorld(l)(0))
	if err != nil {
		l.WithError(err).Errorf("Unable to identify the last world a account was logged into")
	}
}

func announceServerList(l logrus.FieldLogger, ws []world.Model, ms *session.Model) {
	for _, x := range ws {
		err := ms.Announce(writer.WriteServerListEntry(l)(x.Id(), x.Name(), x.Flag(), x.EventMessage(), x.ChannelLoad()))
		if err != nil {
			l.WithError(err).Errorf("Unable to write server list entry for %d", x.Id())
		}
	}
	err := ms.Announce(writer.WriteServerListEnd(l))
	if err != nil {
		l.WithError(err).Errorf("Unable to complete writing the server list")
	}
}
