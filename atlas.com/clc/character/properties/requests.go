package properties

import (
	"atlas-clc/rest/requests"
	"fmt"
)

const (
	CharactersServicePrefix     string = "/ms/cos/"
	CharactersService                  = requests.BaseRequest + CharactersServicePrefix
	CharactersResource                 = CharactersService + "characters/"
	CharactersByName                   = CharactersResource + "?name=%s"
	CharactersForAccountByWorld        = CharactersResource + "?accountId=%d&worldId=%d"
	CharactersById                     = CharactersResource + "%d"
)

func requestPropertiesByName(name string) requests.Request[Attributes] {
	return requests.MakeGetRequest[Attributes](fmt.Sprintf(CharactersByName, name))
}

func requestPropertiesByAccountAndWorld(accountId uint32, worldId byte) requests.Request[Attributes] {
	return requests.MakeGetRequest[Attributes](fmt.Sprintf(CharactersForAccountByWorld, accountId, worldId))
}

func requestPropertiesById(characterId uint32) requests.Request[Attributes] {
	return requests.MakeGetRequest[Attributes](fmt.Sprintf(CharactersById, characterId))
}
