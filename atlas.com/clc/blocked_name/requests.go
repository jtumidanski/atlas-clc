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

func requestByName(name string) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(BlockedNamesByName, name))
}
