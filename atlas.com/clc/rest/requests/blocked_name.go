package requests

import (
	"atlas-clc/rest/attributes"
	"fmt"
)

const (
	BlockedNameResource = BaseRequest + CharactersServicePrefix + "blockedNames/"
	BlockedNamesByName  = BlockedNameResource + "?name=%s"
)

func GetBlockedNamesByName(name string) (*attributes.BlockedNameDataContainer, error) {
	ar := &attributes.BlockedNameDataContainer{}
	err := Get(fmt.Sprintf(BlockedNamesByName, name), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
