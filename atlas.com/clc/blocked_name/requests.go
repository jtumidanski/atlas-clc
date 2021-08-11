package blocked_name

import (
	"atlas-clc/rest/requests"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	BlockedNamesServicePrefix string = "/ms/bns/"
	BlockedNamesService              = requests.BaseRequest + BlockedNamesServicePrefix
	BlockedNameResource              = BlockedNamesService + "names"
	BlockedNamesByName               = BlockedNameResource + "?name=%s"
)

func requestByName(l logrus.FieldLogger) func(name string) (*dataContainer, error) {
	return func(name string) (*dataContainer, error) {
		ar := &dataContainer{}
		err := requests.Get(l)(fmt.Sprintf(BlockedNamesByName, name), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}
