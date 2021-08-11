package world

import (
	"atlas-clc/rest/requests"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	ServicePrefix  string = "/ms/wrg/"
	Service               = requests.BaseRequest + ServicePrefix
	WorldsResource        = Service + "worlds/"
	WorldsById            = WorldsResource + "%d"
)

func requestWorlds(l logrus.FieldLogger) (*dataContainer, error) {
	r := &dataContainer{}
	err := requests.Get(l)(WorldsResource, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func requestWorld(l logrus.FieldLogger) func(worldId byte) (*dataContainer, error) {
	return func(worldId byte) (*dataContainer, error) {
		r := &dataContainer{}
		err := requests.Get(l)(fmt.Sprintf(WorldsById, worldId), r)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}
