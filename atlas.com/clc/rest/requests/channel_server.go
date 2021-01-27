package requests

import (
	"atlas-clc/rest/attributes"
	"fmt"
	"log"
)

const (
	ChannelServersResource            = WorldRegistryService + "channelServers/"
	ChannelServersByWorld             = ChannelServersResource + "?world=%d"
)

func GetChannels(l *log.Logger) (*attributes.ChannelServerDataContainer, error) {
	r := &attributes.ChannelServerDataContainer{}
	err := Get(l, ChannelServersResource, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetChannelsForWorld(l *log.Logger, worldId byte) (*attributes.ChannelServerDataContainer, error) {
	r := &attributes.ChannelServerDataContainer{}
	err := Get(l, fmt.Sprintf(ChannelServersByWorld, worldId), r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
