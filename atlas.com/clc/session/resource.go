package session

import (
	"atlas-clc/json"
	"atlas-clc/socket/response/writer"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
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
		sessionId := getSessionId(l, r)
		errorId := getErrorId(l, r)

		ses := GetRegistry().Get(sessionId)
		if ses == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
		err := ses.Announce(writer.WriteLoginFailed(l)(errorId))
		if err != nil {
			l.WithError(err).Errorf("Unable to issue login failed due to reason %d", errorId)
		}
	}
}

func getSessionId(l logrus.FieldLogger, r *http.Request) int {
	vars := mux.Vars(r)
	value, err := strconv.Atoi(vars["sessionId"])
	if err != nil {
		l.Println("Error parsing worldId as integer")
		return 0
	}
	return value
}

func getErrorId(l logrus.FieldLogger, r *http.Request) byte {
	vars := mux.Vars(r)
	value, err := strconv.Atoi(vars["errorId"])
	if err != nil {
		l.Println("Error parsing worldId as integer")
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
