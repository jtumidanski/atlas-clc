package account

import (
	"atlas-clc/rest/requests"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	AccountsServicePrefix string = "/ms/aos/"
	AccountsService              = requests.BaseRequest + AccountsServicePrefix
	AccountsResource             = AccountsService + "accounts/"
	AccountsByName               = AccountsResource + "?name=%s"
	AccountsById                 = AccountsResource + "%d"
)

type Request func(l logrus.FieldLogger, span opentracing.Span) (*dataContainer, error)

func makeRequest(url string) Request {
	return func(l logrus.FieldLogger, span opentracing.Span) (*dataContainer, error) {
		ar := &dataContainer{}
		err := requests.Get(l, span)(url, ar)
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
