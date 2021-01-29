package resources

import (
	"atlas-clc/registries"
	"atlas-clc/rest/attributes"
	"atlas-clc/mapleSession"
	"atlas-clc/socket/response/writer"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type SessionListDataContainer struct {
	Data []SessionData `json:"data"`
}

type SessionData struct {
	Id         string            `json:"id"`
	Type       string            `json:"type"`
	Attributes SessionAttributes `json:"attributes"`
}

type SessionAttributes struct {
	AccountId int  `json:"accountId"`
	WorldId   byte `json:"worldId"`
	ChannelId byte `json:"channelId"`
}

type SessionResource struct {
	l *log.Logger
}

func NewSessionResource(l *log.Logger) *SessionResource {
	return &SessionResource{l}
}

func (s *SessionResource) GetSessions(rw http.ResponseWriter, _ *http.Request) {
	ss := registries.GetSessionRegistry().GetAll()

	var response SessionListDataContainer
	response.Data = make([]SessionData, 0)
	for _, x := range ss {
		response.Data = append(response.Data, *getSessionObject(x))
	}

	err := attributes.ToJSON(response, rw)
	if err != nil {
		s.l.Println("Error encoding GetSessions response")
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *SessionResource) LoginError(rw http.ResponseWriter, r *http.Request) {
	sessionId := getSessionId(r)
	errorId := getErrorId(r)

	ses := registries.GetSessionRegistry().Get(sessionId)
	if ses != nil {
		ses.Announce(writer.WriteLoginFailed(errorId))
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

func getSessionObject(x mapleSession.MapleSession) *SessionData {
	return &SessionData{
		Id:   strconv.Itoa(x.SessionId()),
		Type: "Session",
		Attributes: SessionAttributes{
			AccountId: x.AccountId(),
			WorldId:   x.WorldId(),
			ChannelId: x.ChannelId(),
		},
	}
}
