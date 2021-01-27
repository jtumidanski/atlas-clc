package requests

import (
	"atlas-clc/rest/attributes"
	"fmt"
	"log"
)

const (
	WorldRegistryServicePrefix string = "/ms/wrg/"
	WorldRegistryService              = BaseRequest + WorldRegistryServicePrefix
	WorldsResource                    = WorldRegistryService + "worlds/"
	WorldsById                        = WorldsResource + "%d"
)

func GetWorlds(l *log.Logger) (*attributes.WorldDataContainer, error) {
	r := &attributes.WorldDataContainer{}
	err := Get(l, WorldsResource, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetWorld(l *log.Logger, worldId byte) (*attributes.WorldDataContainer, error) {
	r := &attributes.WorldDataContainer{}
	err := Get(l, fmt.Sprintf(WorldsById, worldId), r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
