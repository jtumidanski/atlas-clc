package properties

import (
	"atlas-clc/rest/requests"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	CharactersServicePrefix     string = "/ms/cos/"
	CharactersService                  = requests.BaseRequest + CharactersServicePrefix
	CharactersResource                 = CharactersService + "characters/"
	CharactersByName                   = CharactersResource + "?name=%s"
	CharactersForAccountByWorld        = CharactersResource + "?accountId=%d&worldId=%d"
	CharactersById                     = CharactersResource + "%d"
)

type Request func(l logrus.FieldLogger, span opentracing.Span) (*DataContainer, error)

func makeRequest(url string) Request {
	return func(l logrus.FieldLogger, span opentracing.Span) (*DataContainer, error) {
		ar := &DataContainer{}
		err := requests.Get(l, span)(url, ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}

func requestPropertiesByName(name string) Request {
	return makeRequest(fmt.Sprintf(CharactersByName, name))
}

func requestPropertiesByAccountAndWorld(accountId uint32, worldId byte) Request {
	return makeRequest(fmt.Sprintf(CharactersForAccountByWorld, accountId, worldId))
}

func requestPropertiesById(characterId uint32) Request {
	return makeRequest(fmt.Sprintf(CharactersById, characterId))
}