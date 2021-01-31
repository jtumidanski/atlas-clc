package requests

import (
	"atlas-clc/rest/attributes"
	"fmt"
)

const (
	WorldRegistryServicePrefix string = "/ms/wrg/"
	WorldRegistryService              = BaseRequest + WorldRegistryServicePrefix
	WorldsResource                    = WorldRegistryService + "worlds/"
	WorldsById                        = WorldsResource + "%d"
)

func GetWorlds() (*attributes.WorldDataContainer, error) {
	r := &attributes.WorldDataContainer{}
	err := Get(WorldsResource, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetWorld(worldId byte) (*attributes.WorldDataContainer, error) {
	r := &attributes.WorldDataContainer{}
	err := Get(fmt.Sprintf(WorldsById, worldId), r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
