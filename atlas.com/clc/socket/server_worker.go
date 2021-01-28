package socket

import (
	"atlas-clc/registries"
	"atlas-clc/rest/requests"
	"atlas-clc/sessions"
	"log"
)

func Disconnect(l *log.Logger, s *sessions.Session) {
	requests.CreateLogout(s.AccountId())
	registries.GetSessionRegistry().Remove(s.SessionId())
}
