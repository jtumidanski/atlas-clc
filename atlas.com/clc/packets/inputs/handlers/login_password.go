package handlers

import (
	"atlas-clc/packets/inputs"
	"atlas-clc/packets/inputs/models"
	"atlas-clc/packets/inputs/readers"
	"atlas-clc/packets/outputs/constants"
	"atlas-clc/packets/outputs/writers"
	"atlas-clc/processors"
	"atlas-clc/registries"
	"atlas-clc/rest/attributes"
	"atlas-clc/rest/requests"
	"atlas-clc/sessions"
	"log"
	"net/http"
	"strconv"
)

type LoginPasswordHandler struct {
}

func (h *LoginPasswordHandler) Handle(l *log.Logger, sessionId int, r *inputs.Reader) {
	s := registries.GetSessionRegistry().GetSession(sessionId)
	p := readers.ReadLoginPassword(r)
	h.handle(l, s, p)
}

func (h *LoginPasswordHandler) handle(l *log.Logger, s *sessions.Session, p *models.LoginPassword) {
	ip := s.GetRemoteAddress().String()
	r, err := requests.CreateLogin(l, s.SessionId(), p.Login(), p.Password(), ip)
	if err != nil {
		h.announceSystemError(s)
		return
	}

	if r.StatusCode != http.StatusNoContent {
		eb := &attributes.ErrorListDataContainer{}
		err = requests.ProcessErrorResponse(l, r, eb)
		if err != nil {
			h.announceSystemError(s)
			return
		}

		if len(eb.Errors) > 0 {
			h.processFirstError(l, s, eb.Errors[0])
			return
		}

		h.announceSystemError(s)
		return
	}

	h.authorizeSuccess(l, s, p.Login())
}

func (h *LoginPasswordHandler) authorizeSuccess(l *log.Logger, s *sessions.Session, name string) {
	a, err := processors.GetAccountByName(l, name)
	if err == nil {
		s.SetAccountId(a.Id())
		s.Announce(writers.WriteAuthSuccess(a.Id(), a.Name(), a.Gender(), a.PIC()))
	}
}

func (h *LoginPasswordHandler) announceSystemError(s *sessions.Session) {
	s.Announce(writers.WriteLoginFailed(constants.SystemError))
}

func (h LoginPasswordHandler) processFirstError(l *log.Logger, s *sessions.Session, data attributes.ErrorData) {
	r := constants.GetLoginFailedReason(data.Code)
	if r == constants.DeletedOrBlocked {
		if data.Detail == "" {
			s.Announce(writers.WriteLoginFailed(constants.DeletedOrBlocked))
			return
		}

		reason := data.Meta["reason"]
		rc, err := strconv.ParseUint(reason, 10, 8)
		if err != nil {
			s.Announce(writers.WriteLoginFailed(constants.SystemError))
			return
		}

		if tb, ok := data.Meta["tempBan"]; ok {
			until, err := strconv.ParseUint(tb, 10, 64)
			if err != nil {
				s.Announce(writers.WriteLoginFailed(constants.SystemError))
				return
			}
			s.Announce(writers.WriteTemporaryBan(until, byte(rc)))
			return
		}
		s.Announce(writers.WritePermanentBan(byte(rc)))
		return
	}
	s.Announce(writers.WriteLoginFailed(r))
}
