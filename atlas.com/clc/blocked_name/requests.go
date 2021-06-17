package blocked_name

import (
	"atlas-clc/rest/requests"
	"fmt"
)

const (
	BlockedNamesServicePrefix string = "/ms/bns/"
	BlockedNamesService              = requests.BaseRequest + BlockedNamesServicePrefix
	BlockedNameResource              = BlockedNamesService + "names"
	BlockedNamesByName               = BlockedNameResource + "?name=%s"
)

func GetBlockedNamesByName(name string) (*BlockedNameDataContainer, error) {
	ar := &BlockedNameDataContainer{}
	err := requests.Get(fmt.Sprintf(BlockedNamesByName, name), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
