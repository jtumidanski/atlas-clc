package handler

import (
	"atlas-clc/account"
	"atlas-clc/login"
	"atlas-clc/rest/requests"
	"atlas-clc/rest/resources"
	"atlas-clc/session"
	"atlas-clc/socket/response/writer"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/sirupsen/logrus"
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

func (h *LoginHandler) IsValid(_ logrus.FieldLogger, _ *session.MapleSession) bool {
	return true
}

func (h *LoginHandler) HandleRequest(_ logrus.FieldLogger, ms *session.MapleSession, r *request.RequestReader) {
	p := ReadLoginRequest(r)

	ip := (*ms).GetRemoteAddress().String()
	resp, err := login.CreateLogin((*ms).SessionId(), p.Login(), p.Password(), ip)
	if err != nil {
		h.announceSystemError(ms)
		return
	}

	if resp.StatusCode != http.StatusNoContent {
		eb := &resources.ErrorListDataContainer{}
		err = requests.ProcessErrorResponse(resp, eb)
		if err != nil {
			h.announceSystemError(ms)
			return
		}

		if len(eb.Errors) > 0 {
			h.processFirstError(ms, eb.Errors[0])
			return
		}

		h.announceSystemError(ms)
		return
	}

	h.authorizeSuccess(ms, p.Login())
}

func (h *LoginHandler) authorizeSuccess(ms *session.MapleSession, name string) {
	a, err := account.GetByName(name)
	if err == nil {
		(*ms).SetAccountId(a.Id())
		(*ms).Announce(writer.WriteAuthSuccess(a.Id(), a.Name(), a.Gender(), a.PIC()))
	}
}

func (h *LoginHandler) announceSystemError(ms *session.MapleSession) {
	(*ms).Announce(writer.WriteLoginFailed(SystemError))
}

func (h LoginHandler) processFirstError(ms *session.MapleSession, data resources.ErrorData) {
	r := GetLoginFailedReason(data.Code)
	if r == DeletedOrBlocked {
		if data.Detail == "" {
			(*ms).Announce(writer.WriteLoginFailed(DeletedOrBlocked))
			return
		}

		reason := data.Meta["reason"]
		rc, err := strconv.ParseUint(reason, 10, 8)
		if err != nil {
			(*ms).Announce(writer.WriteLoginFailed(SystemError))
			return
		}

		if tb, ok := data.Meta["tempBan"]; ok {
			until, err := strconv.ParseUint(tb, 10, 64)
			if err != nil {
				(*ms).Announce(writer.WriteLoginFailed(SystemError))
				return
			}
			(*ms).Announce(writer.WriteTemporaryBan(until, byte(rc)))
			return
		}
		(*ms).Announce(writer.WritePermanentBan())
		return
	}
	(*ms).Announce(writer.WriteLoginFailed(r))
}
