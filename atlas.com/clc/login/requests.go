package login

import (
	"atlas-clc/account"
	"atlas-clc/rest/requests"
	"net/http"
)

const (
	LoginsResource = account.AccountsService + "logins/"
)

func CreateLogin(sessionId int, name string, password string, ipAddress string) (r *http.Response, err error) {
	i := LoginInputDataContainer{
		Data: LoginData{
			Id:   "0",
			Type: "com.atlas.aos.attribute.LoginAttributes",
			Attributes: LoginAttributes{
				SessionId: sessionId,
				Name:      name,
				Password:  password,
				IpAddress: ipAddress,
				State:     0,
			},
		},
	}

	return requests.Post(LoginsResource, i)
}
