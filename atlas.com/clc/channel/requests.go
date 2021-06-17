package channel

import (
	"atlas-clc/rest/requests"
	"fmt"
)

const (
	ChannelRegistryServicePrefix string = "/ms/wrg/"
	ChannelRegistryService              = requests.BaseRequest + ChannelRegistryServicePrefix
	ChannelServersResource              = ChannelRegistryService + "channelServers/"
	ChannelServersByWorld               = ChannelServersResource + "?world=%d"
)

func requestChannels() (*ChannelServerDataContainer, error) {
	r := &ChannelServerDataContainer{}
	err := requests.Get(ChannelServersResource, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func requestChannelsForWorld(worldId byte) (*ChannelServerDataContainer, error) {
	r := &ChannelServerDataContainer{}
	err := requests.Get(fmt.Sprintf(ChannelServersByWorld, worldId), r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
