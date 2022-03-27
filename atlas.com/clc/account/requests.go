package account

import (
	"atlas-clc/rest/requests"
	"fmt"
)

const (
	AccountsServicePrefix string = "/ms/aos/"
	AccountsService              = requests.BaseRequest + AccountsServicePrefix
	AccountsResource             = AccountsService + "accounts/"
	AccountsByName               = AccountsResource + "?name=%s"
	AccountsById                 = AccountsResource + "%d"
)

func requestAccountByName(name string) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(AccountsByName, name))
}

func requestAccountById(id uint32) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(AccountsById, id))
}
