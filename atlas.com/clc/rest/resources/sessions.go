package resources

import (
   "atlas-clc/packets/outputs/writers"
   "atlas-clc/registries"
   "atlas-clc/rest/attributes"
   "atlas-clc/sessions"
   "github.com/gorilla/mux"
   "log"
   "net/http"
   "strconv"
)

type Session struct {
   l *log.Logger
}

func NewSession(l *log.Logger) *Session {
   return &Session{l}
}

func (s *Session) GetSessions(rw http.ResponseWriter, r *http.Request) {
   ss := registries.GetSessionRegistry().GetSessions()

   var response attributes.SessionListDataContainer
   response.Data = make([]attributes.SessionData, 0)
   for _, x := range ss {
      response.Data = append(response.Data, *getSessionObject(x))
   }

   err := attributes.ToJSON(response, rw)
   if err != nil {
      s.l.Println("Error encoding GetSessions response")
      rw.WriteHeader(http.StatusInternalServerError)
   }
}

func (s *Session) LoginError(rw http.ResponseWriter, r *http.Request) {
   sessionId := getSessionId(r)
   errorId := getErrorId(r)

   ses := registries.GetSessionRegistry().GetSession(sessionId)
   if ses != nil {
      ses.Announce(writers.WriteLoginFailed(errorId))
      rw.WriteHeader(http.StatusNoContent)
   } else {
      rw.WriteHeader(http.StatusNotFound)
   }
}

func getSessionId(r *http.Request) int {
   vars := mux.Vars(r)
   value, err := strconv.Atoi(vars["sessionId"])
   if err != nil {
      log.Println("Error parsing worldId as integer")
      return 0
   }
   return value
}

func getErrorId(r *http.Request) byte {
   vars := mux.Vars(r)
   value, err := strconv.Atoi(vars["errorId"])
   if err != nil {
      log.Println("Error parsing worldId as integer")
      return 0
   }
   return byte(value)
}

func getSessionObject(x sessions.Session) *attributes.SessionData {
   return &attributes.SessionData{
      Id:   strconv.Itoa(x.SessionId()),
      Type: "Session",
      Attributes: attributes.SessionAttributes{
         AccountId: x.AccountId(),
      },
   }
}
