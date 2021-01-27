package requests

import (
	"atlas-clc/rest/attributes"
	"fmt"
	"log"
)

func GetWorlds(l *log.Logger) (*attributes.WorldListDataContainer, error) {
	r := &attributes.WorldListDataContainer{}
	err := Get(l, "http://atlas-nginx:80/ms/wrg/worlds/", r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetWorld(l *log.Logger, worldId byte) (*attributes.WorldDataContainer, error) {
	r := &attributes.WorldDataContainer{}
	err := Get(l, fmt.Sprintf("http://atlas-nginx:80/ms/wrg/worlds/%d", worldId), r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
