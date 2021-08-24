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

type Request func(l logrus.FieldLogger) (*dataContainer, error)

func makeRequest(url string) Request {
	return func(l logrus.FieldLogger) (*dataContainer, error) {
		ar := &dataContainer{}
		err := requests.Get(l)(url, ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}

func requestAccountByName(name string) Request {
	return makeRequest(fmt.Sprintf(AccountsByName, name))
}

func requestAccountById(id uint32) Request {
	return makeRequest(fmt.Sprintf(AccountsById, id))
}
