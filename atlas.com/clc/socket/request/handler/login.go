package handler

import (
	"atlas-clc/processors"
	"atlas-clc/rest/attributes"
	"atlas-clc/rest/requests"
	"atlas-clc/sessions"
	"atlas-clc/socket/request"
	"atlas-clc/socket/response/writer"
	"log"
	"net/http"
	"strconv"
)

const OpCodeLogin uint16 = 0x01

type LoginRequest struct {
	login    string
	password string
	hwid     []byte
}

func (l *LoginRequest) Login() string {
	return l.login
}

func (l *LoginRequest) Password() string {
	return l.password
}

func ReadLoginRequest(reader *request.RequestReader) *LoginRequest {
	login := reader.ReadAsciiString()
	password := reader.ReadAsciiString()
	reader.Skip(6)
	hwid := reader.ReadBytes(4)
	return &LoginRequest{
		login:    login,
		password: password,
		hwid:     hwid,
	}
}

type LoginHandler struct {
}

func (h *LoginHandler) IsValid(_ *log.Logger, _ *sessions.Session) bool {
	return true
}

func (h *LoginHandler) HandleRequest(l *log.Logger, s *sessions.Session, r *request.RequestReader) {
	p := ReadLoginRequest(r)

	ip := s.GetRemoteAddress().String()
	resp, err := requests.CreateLogin(s.SessionId(), p.Login(), p.Password(), ip)
	if err != nil {
		h.announceSystemError(s)
		return
	}

	if resp.StatusCode != http.StatusNoContent {
		eb := &attributes.ErrorListDataContainer{}
		err = requests.ProcessErrorResponse(resp, eb)
		if err != nil {
			h.announceSystemError(s)
			return
		}

		if len(eb.Errors) > 0 {
			h.processFirstError(s, eb.Errors[0])
			return
		}

		h.announceSystemError(s)
		return
	}

	h.authorizeSuccess(l, s, p.Login())
}

func (h *LoginHandler) authorizeSuccess(l *log.Logger, s *sessions.Session, name string) {
	a, err := processors.GetAccountByName(name)
	if err == nil {
		s.SetAccountId(a.Id())
		s.Announce(writer.WriteAuthSuccess(a.Id(), a.Name(), a.Gender(), a.PIC()))
	}
}

func (h *LoginHandler) announceSystemError(s *sessions.Session) {
	s.Announce(writer.WriteLoginFailed(SystemError))
}

func (h LoginHandler) processFirstError(s *sessions.Session, data attributes.ErrorData) {
	r := GetLoginFailedReason(data.Code)
	if r == DeletedOrBlocked {
		if data.Detail == "" {
			s.Announce(writer.WriteLoginFailed(DeletedOrBlocked))
			return
		}

		reason := data.Meta["reason"]
		rc, err := strconv.ParseUint(reason, 10, 8)
		if err != nil {
			s.Announce(writer.WriteLoginFailed(SystemError))
			return
		}

		if tb, ok := data.Meta["tempBan"]; ok {
			until, err := strconv.ParseUint(tb, 10, 64)
			if err != nil {
				s.Announce(writer.WriteLoginFailed(SystemError))
				return
			}
			s.Announce(writer.WriteTemporaryBan(until, byte(rc)))
			return
		}
		s.Announce(writer.WritePermanentBan())
		return
	}
	s.Announce(writer.WriteLoginFailed(r))
}
