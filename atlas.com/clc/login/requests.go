package login

import (
	"atlas-clc/account"
	"atlas-clc/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	LoginsResource = account.AccountsService + "logins/"
)

func CreateLogin(l logrus.FieldLogger, span opentracing.Span) func(sessionId uint32, name string, password string, ipAddress string) (*requests.ErrorListDataContainer, error) {
	return func(sessionId uint32, name string, password string, ipAddress string) (*requests.ErrorListDataContainer, error) {
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
		resp := &dataContainer{}
		errResp := &requests.ErrorListDataContainer{}
		err := requests.Post(l, span)(LoginsResource, i, resp, errResp)
		if err != nil {
			return nil, err
		}
		return errResp, nil
	}
}
