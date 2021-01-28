package requests

import (
	"atlas-clc/rest/attributes"
	"fmt"
)

const (
	AccountsServicePrefix string = "/ms/aos/"
	AccountsService              = BaseRequest + AccountsServicePrefix
	AccountsResource             = AccountsService + "accounts/"
	AccountsByName               = AccountsResource + "?name=%s"
	AccountsById                 = AccountsResource + "%d"
)

func GetAccountByName(name string) (*attributes.AccountDataContainer, error) {
	ar := &attributes.AccountDataContainer{}
	err := Get(fmt.Sprintf(AccountsByName, name), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func GetAccountById(id int) (*attributes.AccountDataContainer, error) {
	ar := &attributes.AccountDataContainer{}
	err := Get(fmt.Sprintf(AccountsById, id), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
