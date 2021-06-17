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

func requestByName(name string) (*dataContainer, error) {
	ar := &dataContainer{}
	err := requests.Get(fmt.Sprintf(BlockedNamesByName, name), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
