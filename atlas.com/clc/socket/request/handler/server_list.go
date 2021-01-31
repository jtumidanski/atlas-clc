package handler

import (
   "atlas-clc/domain"
   "atlas-clc/mapleSession"
   "atlas-clc/processors"
   "atlas-clc/socket/response/writer"
   "github.com/jtumidanski/atlas-socket/request"
   "log"
)

const OpCodeServerRequest uint16 = 0x0B
const OpCodeServerListReRequest uint16 = 0x04

type ServerListHandler struct {
}

func (h *ServerListHandler) IsValid(l *log.Logger, ms *mapleSession.MapleSession) bool {
   v := processors.IsLoggedIn((*ms).AccountId())
   if !v {
      l.Printf("[ERROR] attempting to process a [ServerListRequest] when the account %d is not logged in.", (*ms).SessionId())
   }
   return v
}

func (h *ServerListHandler) HandleRequest(l *log.Logger, ms *mapleSession.MapleSession, _ *request.RequestReader) {
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
   h.respondToSession(ms, nws)
}

func (h *ServerListHandler) combine(l *log.Logger, ws []domain.World, cls map[int][]domain.ChannelLoad) []domain.World {
   var nws = make([]domain.World, 0)

   for _, x := range ws {
      if val, ok := cls[int(x.Id())]; ok {
         nws = append(nws, x.SetChannelLoad(val))
      } else {
         l.Println("[ERROR] processing world without a channel load")
         nws = append(nws, x)
      }
   }
   return nws
}

func (h *ServerListHandler) respondToSession(ms *mapleSession.MapleSession, ws []domain.World) {
   h.announceServerList(ws, ms)
   h.announceLastWorld(ms)
   h.announceRecommendedWorlds(ws, ms)
}

func (h *ServerListHandler) announceRecommendedWorlds(ws []domain.World, ms *mapleSession.MapleSession) {
   var rs = make([]domain.WorldRecommendation, 0)
   for _, x := range ws {
      if x.Recommended() {
         rs = append(rs, x.Recommendation())
      }
   }
   (*ms).Announce(writer.WriteRecommendedWorlds(rs))
}

func (h *ServerListHandler) announceLastWorld(ms *mapleSession.MapleSession) {
   (*ms).Announce(writer.WriteSelectWorld(0))
}

func (h *ServerListHandler) announceServerList(ws []domain.World, ms *mapleSession.MapleSession) {
   for _, x := range ws {
      (*ms).Announce(writer.WriteServerListEntry(x.Id(), x.Name(), x.Flag(), x.EventMessage(), x.ChannelLoad()))
   }
   (*ms).Announce(writer.WriteServerListEnd())
}
