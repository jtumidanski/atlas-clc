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

func requestWorlds() requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](WorldsResource)
}

func requestWorld(worldId byte) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(WorldsById, worldId))
}
