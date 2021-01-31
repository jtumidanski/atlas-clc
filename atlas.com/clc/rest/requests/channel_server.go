package requests

import (
	"atlas-clc/rest/attributes"
	"fmt"
)

const (
	ChannelServersResource            = WorldRegistryService + "channelServers/"
	ChannelServersByWorld             = ChannelServersResource + "?world=%d"
)

func GetChannels() (*attributes.ChannelServerDataContainer, error) {
	r := &attributes.ChannelServerDataContainer{}
	err := Get(ChannelServersResource, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetChannelsForWorld(worldId byte) (*attributes.ChannelServerDataContainer, error) {
	r := &attributes.ChannelServerDataContainer{}
	err := Get(fmt.Sprintf(ChannelServersByWorld, worldId), r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
