package session

import (
	"atlas-clc/json"
	"atlas-clc/socket/response/writer"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strconv"
)

func HandleGetSessions(l logrus.FieldLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ss := GetRegistry().GetAll()

		var response dataListContainer
		response.Data = make([]dataBody, 0)
		for _, x := range ss {
			response.Data = append(response.Data, *getSessionObject(x))
		}

		err := json.ToJSON(response, w)
		if err != nil {
			l.WithError(err).Errorf("Error encoding GetSessions response")
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func HandleLoginError(l logrus.FieldLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionId := getSessionId(r)
		errorId := getErrorId(r)

		ses := GetRegistry().Get(sessionId)
		if ses != nil {
			ses.Announce(writer.WriteLoginFailed(errorId))
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
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

func getSessionObject(x MapleSession) *dataBody {
	return &dataBody{
		Id:   strconv.Itoa(x.SessionId()),
		Type: "Session",
		Attributes: attributes{
			AccountId: x.AccountId(),
			WorldId:   x.WorldId(),
			ChannelId: x.ChannelId(),
		},
	}
}
