package world

import (
	"atlas-clc/rest/requests"
	"fmt"
)

const (
	ServicePrefix  string = "/ms/wrg/"
	Service               = requests.BaseRequest + ServicePrefix
	WorldsResource        = Service + "worlds/"
	WorldsById            = WorldsResource + "%d"
)

func requestWorlds() (*dataContainer, error) {
	r := &dataContainer{}
	err := requests.Get(WorldsResource, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func requestWorld(worldId byte) (*dataContainer, error) {
	r := &dataContainer{}
	err := requests.Get(fmt.Sprintf(WorldsById, worldId), r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
