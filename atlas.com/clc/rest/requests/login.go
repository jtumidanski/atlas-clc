package requests

import (
	"atlas-clc/rest/attributes"
	"fmt"
	"net/http"
)

const (
	LoginsResource = AccountsService + "logins/"
	LoginsById     = LoginsResource + "%d"
)

func CreateLogin(sessionId int, name string, password string, ipAddress string) (r *http.Response, err error) {
	i := attributes.LoginInputDataContainer{
		Data: attributes.LoginData{
			Id:   "0",
			Type: "com.atlas.aos.attribute.LoginAttributes",
			Attributes: attributes.LoginAttributes{
				SessionId: sessionId,
				Name:      name,
				Password:  password,
				IpAddress: ipAddress,
				State:     0,
			},
		},
	}

	return Post(LoginsResource, i)
}

func CreateLogout(accountId int) {
	_, _ = Delete(fmt.Sprintf(LoginsById, accountId))
}
