package handler

import (
	"atlas-clc/account"
	"atlas-clc/login"
	"atlas-clc/model"
	"atlas-clc/rest/requests"
	"atlas-clc/session"
	"atlas-clc/socket/response/writer"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
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

func HandleLoginRequest(l logrus.FieldLogger, span opentracing.Span) func(s session.Model, r *request.RequestReader) {
	return func(s session.Model, r *request.RequestReader) {
		p := ReadLoginRequest(r)
		ip := s.GetRemoteAddress().String()
		respErr, err := login.CreateLogin(l, span)(s.SessionId(), p.Name(), p.Password(), ip)
		if err != nil {
			announceError(l, span)(s, SystemError)
			return
		}

		if len(respErr.Errors) > 0 {
			processFirstError(l, span)(s, respErr.Errors[0])
			return
		}

		account.ForAccountByName(l, span)(p.Name(), issueSuccess(l, s))
	}
}

func issueSuccess(l logrus.FieldLogger, ms session.Model) model.Operator[account.Model] {
	return func(a account.Model) error {
		ms = session.SetAccountId(a.Id())(ms.SessionId())
		err := session.Announce(writer.WriteAuthSuccess(l)(a.Id(), a.Name(), a.Gender(), a.PIC()))(ms)
		if err != nil {
			l.WithError(err).Errorf("Unable to show successful authorization for account %d", a.Id())
		}
		return err
	}
}

func announceError(l logrus.FieldLogger, _ opentracing.Span) func(s session.Model, reason byte) {
	return func(s session.Model, reason byte) {
		err := session.Announce(writer.WriteLoginFailed(l)(reason))(s)
		if err != nil {
			l.WithError(err).WithField("reason", reason).Errorf("Unable to identify to character that login has failed.")
		}
	}
}

func processFirstError(l logrus.FieldLogger, span opentracing.Span) func(s session.Model, data requests.ErrorData) {
	return func(s session.Model, data requests.ErrorData) {
		r := GetLoginFailedReason(data.Code)
		if r == DeletedOrBlocked {
			if data.Detail == "" {
				announceError(l, span)(s, DeletedOrBlocked)
				return
			}

			reason := data.Meta["reason"]
			rc, err := strconv.ParseUint(reason, 10, 8)
			if err != nil {
				announceError(l, span)(s, SystemError)
				return
			}

			if tb, ok := data.Meta["tempBan"]; ok {
				until, err := strconv.ParseUint(tb, 10, 64)
				if err != nil {
					announceError(l, span)(s, SystemError)
					return
				}
				err = session.Announce(writer.WriteTemporaryBan(l)(until, byte(rc)))(s)
				if err != nil {
					l.WithError(err).Errorf("Unable to issue login failed due to temporary ban")
				}
				return
			}
			err = session.Announce(writer.WritePermanentBan(l))(s)
			if err != nil {
				l.WithError(err).Errorf("Unable to issue login failed due to permanent ban")
			}
			return
		}
		announceError(l, span)(s, r)
	}
}
