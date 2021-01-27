package requests

import (
	"atlas-clc/rest/attributes"
	"fmt"
	"log"
)

const (
	BlockedNameResource = BaseRequest + CharactersServicePrefix + "blockedNames/"
	BlockedNamesByName  = BlockedNameResource + "?name=%s"
)

func GetBlockedNamesByName(l *log.Logger, name string) (*attributes.BlockedNameDataContainer, error) {
	ar := &attributes.BlockedNameDataContainer{}
	err := Get(l, fmt.Sprintf(BlockedNamesByName, name), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
