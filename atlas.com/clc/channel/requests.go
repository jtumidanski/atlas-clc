package channel

import (
	"atlas-clc/rest/requests"
	"fmt"
)

const (
	ServicePrefix string = "/ms/wrg/"
	Service              = requests.BaseRequest + ServicePrefix
	Resource             = Service + "channelServers/"
	ByWorld              = Resource + "?world=%d"
)

func requestChannels() requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](Resource)
}

func requestChannelsForWorld(worldId byte) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(ByWorld, worldId))
}
