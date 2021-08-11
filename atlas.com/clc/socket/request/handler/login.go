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
	name     string
	password string
	hwid     []byte
}

func (l *LoginRequest) Name() string {
	return l.name
}

func (l *LoginRequest) Password() string {
	return l.password
}

func ReadLoginRequest(reader *request.RequestReader) *LoginRequest {
	name := reader.ReadAsciiString()
	password := reader.ReadAsciiString()
	reader.Skip(6)
	hwid := reader.ReadBytes(4)
	return &LoginRequest{
		name:     name,
		password: password,
		hwid:     hwid,
	}
}

func HandleLoginRequest(l logrus.FieldLogger, ms *session.Model, r *request.RequestReader) {
	p := ReadLoginRequest(r)

	ip := ms.GetRemoteAddress().String()
	resp, err := login.CreateLogin(ms.SessionId(), p.Name(), p.Password(), ip)
	if err != nil {
		announceSystemError(l, ms)
		return
	}

	if resp.StatusCode != http.StatusNoContent {
		eb := &resources.ErrorListDataContainer{}
		err = requests.ProcessErrorResponse(resp, eb)
		if err != nil {
			announceSystemError(l, ms)
			return
		}

		if len(eb.Errors) > 0 {
			processFirstError(l, ms, eb.Errors[0])
			return
		}

		announceSystemError(l, ms)
		return
	}

	authorizeSuccess(l, ms, p.Name())
}

func authorizeSuccess(l logrus.FieldLogger, ms *session.Model, name string) {
	a, err := account.GetByName(l)(name)
	if err == nil {
		ms.SetAccountId(a.Id())
		err = ms.Announce(writer.WriteAuthSuccess(l)(a.Id(), a.Name(), a.Gender(), a.PIC()))
		if err != nil {
			l.WithError(err).Errorf("Unable to show successful authorization for account %d", a.Id())
		}
	}
}

func announceSystemError(l logrus.FieldLogger, ms *session.Model) {
	err := ms.Announce(writer.WriteLoginFailed(l)(SystemError))
	if err != nil {
		l.WithError(err).Errorf("Unable to identify that login has failed")
	}
}

func processFirstError(l logrus.FieldLogger, ms *session.Model, data resources.ErrorData) {
	r := GetLoginFailedReason(data.Code)
	if r == DeletedOrBlocked {
		if data.Detail == "" {
			err := ms.Announce(writer.WriteLoginFailed(l)(DeletedOrBlocked))
			if err != nil {
				l.WithError(err).Errorf("Unable to issue login failed due to account being deleted or blocked")
			}
			return
		}

		reason := data.Meta["reason"]
		rc, err := strconv.ParseUint(reason, 10, 8)
		if err != nil {
			err = ms.Announce(writer.WriteLoginFailed(l)(SystemError))
			if err != nil {
				l.WithError(err).Errorf("Unable to issue login failed due to system error")
			}
			return
		}

		if tb, ok := data.Meta["tempBan"]; ok {
			until, err := strconv.ParseUint(tb, 10, 64)
			if err != nil {
				err = ms.Announce(writer.WriteLoginFailed(l)(SystemError))
				if err != nil {
					l.WithError(err).Errorf("Unable to issue login failed due to system error")
				}
				return
			}
			err = ms.Announce(writer.WriteTemporaryBan(l)(until, byte(rc)))
			if err != nil {
				l.WithError(err).Errorf("Unable to issue login failed due to temporary ban")
			}
			return
		}
		err = ms.Announce(writer.WritePermanentBan(l))
		if err != nil {
			l.WithError(err).Errorf("Unable to issue login failed due to permanent ban")
		}
		return
	}
	err := ms.Announce(writer.WriteLoginFailed(l)(r))
	if err != nil {
		l.WithError(err).Errorf("Unable to issue login failed due to reason %d", r)
	}
}
