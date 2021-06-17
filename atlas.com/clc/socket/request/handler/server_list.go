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
   ws, err := world.GetWorlds()
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
   h.respondToSession(ms, nws)
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

func (h *ServerListHandler) respondToSession(ms *session.MapleSession, ws []world.Model) {
   h.announceServerList(ws, ms)
   h.announceLastWorld(ms)
   h.announceRecommendedWorlds(ws, ms)
}

func (h *ServerListHandler) announceRecommendedWorlds(ws []world.Model, ms *session.MapleSession) {
   var rs = make([]world.Recommendation, 0)
   for _, x := range ws {
      if x.Recommended() {
         rs = append(rs, x.Recommendation())
      }
   }
   (*ms).Announce(writer.WriteRecommendedWorlds(rs))
}

func (h *ServerListHandler) announceLastWorld(ms *session.MapleSession) {
   (*ms).Announce(writer.WriteSelectWorld(0))
}

func (h *ServerListHandler) announceServerList(ws []world.Model, ms *session.MapleSession) {
   for _, x := range ws {
      (*ms).Announce(writer.WriteServerListEntry(x.Id(), x.Name(), x.Flag(), x.EventMessage(), x.ChannelLoad()))
   }
   (*ms).Announce(writer.WriteServerListEnd())
}
