package world

import (
	"atlas-clc/rest/requests"
	"fmt"
)

const (
	WorldRegistryServicePrefix string = "/ms/wrg/"
	WorldRegistryService              = requests.BaseRequest + WorldRegistryServicePrefix
	WorldsResource                    = WorldRegistryService + "worlds/"
	WorldsById                        = WorldsResource + "%d"
)

func requestWorlds() (*WorldDataContainer, error) {
	r := &WorldDataContainer{}
	err := requests.Get(WorldsResource, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func requestWorld(worldId byte) (*WorldDataContainer, error) {
	r := &WorldDataContainer{}
	err := requests.Get(fmt.Sprintf(WorldsById, worldId), r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
