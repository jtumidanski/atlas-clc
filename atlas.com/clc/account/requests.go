package account

import (
	"atlas-clc/rest/requests"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	AccountsServicePrefix string = "/ms/aos/"
	AccountsService              = requests.BaseRequest + AccountsServicePrefix
	AccountsResource             = AccountsService + "accounts/"
	AccountsByName               = AccountsResource + "?name=%s"
	AccountsById                 = AccountsResource + "%d"
)

func requestAccountByName(l logrus.FieldLogger) func(name string) (*dataContainer, error) {
	return func(name string) (*dataContainer, error) {
		ar := &dataContainer{}
		err := requests.Get(l)(fmt.Sprintf(AccountsByName, name), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}

func requestAccountById(l logrus.FieldLogger) func(id uint32) (*dataContainer, error) {
	return func(id uint32) (*dataContainer, error) {
		ar := &dataContainer{}
		err := requests.Get(l)(fmt.Sprintf(AccountsById, id), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}
