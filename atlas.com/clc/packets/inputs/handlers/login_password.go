package handlers

import (
   "atlas-clc/packets"
   "atlas-clc/packets/inputs/models"
   "atlas-clc/packets/inputs/readers"
   "atlas-clc/registries"
   "atlas-clc/rest/requests"
   "atlas-clc/sessions"
   "log"
   "net/http"
)

func HandleLoginPassword(sessionId int, r *packets.Reader, l *log.Logger) {
   s := registries.GetSessionRegistry().GetSession(sessionId)
   p := readers.ReadLoginPassword(r)
   handle(s, p, l)
}

func handle(s *sessions.Session, p *models.LoginPassword, l *log.Logger) {
   ip := s.GetRemoteAddress().String()
   sc, r, e := requests.CreateLogin(l, s.SessionId(), p.Login(), p.Password(), ip)
   if sc == http.StatusNoContent {
      // get account by name
      // store account id in session
      // write auth success
   } else if len(e.Errors) > 0 {

   } else {
      // total failure, no idea what happened
   }
}
