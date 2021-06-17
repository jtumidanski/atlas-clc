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

func requestAccountByName(name string) (*dataContainer, error) {
	ar := &dataContainer{}
	err := requests.Get(fmt.Sprintf(AccountsByName, name), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func requestAccountById(id uint32) (*dataContainer, error) {
	ar := &dataContainer{}
	err := requests.Get(fmt.Sprintf(AccountsById, id), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
