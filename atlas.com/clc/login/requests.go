package login

import (
	"atlas-clc/account"
	"atlas-clc/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	LoginsResource = account.AccountsService + "logins/"
)

func CreateLogin(l logrus.FieldLogger, span opentracing.Span) func(sessionId uint32, name string, password string, ipAddress string) (r *http.Response, err error) {
	return func(sessionId uint32, name string, password string, ipAddress string) (r *http.Response, err error) {
		i := inputDataContainer{
			Data: dataBody{
				Id:   "0",
				Type: "com.atlas.aos.attribute.LoginAttributes",
				Attributes: attributes{
					SessionId: sessionId,
					Name:      name,
					Password:  password,
					IpAddress: ipAddress,
					State:     0,
				},
			},
		}

		return requests.Post(l, span)(LoginsResource, i)
	}
}
