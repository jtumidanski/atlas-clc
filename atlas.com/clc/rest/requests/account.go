package requests

import (
	"atlas-clc/rest/attributes"
	"fmt"
	"log"
)

const (
	AccountsServicePrefix string = "/ms/aos/"
	AccountsService              = BaseRequest + AccountsServicePrefix
	AccountsResource             = AccountsService + "accounts/"
	AccountsByName               = AccountsResource + "?name=%s"
	AccountsById                 = AccountsResource + "%d"
)

func GetAccountByName(l *log.Logger, name string) (*attributes.AccountDataContainer, error) {
	ar := &attributes.AccountDataContainer{}
	err := Get(l, fmt.Sprintf(AccountsByName, name), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func GetAccountById(l *log.Logger, id int) (*attributes.AccountDataContainer, error) {
	ar := &attributes.AccountDataContainer{}
	err := Get(l, fmt.Sprintf(AccountsById, id), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
