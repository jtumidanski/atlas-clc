package requests

import (
	"atlas-clc/rest/attributes"
	"fmt"
	"log"
)

func GetChannels(l *log.Logger) (*attributes.ChannelServerListDataContainer, error) {
	r := &attributes.ChannelServerListDataContainer{}
	err := Get(l, "http://atlas-nginx:80/ms/wrg/channelServers", r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetChannelsForWorld(l *log.Logger, worldId byte) (*attributes.ChannelServerListDataContainer, error) {
	r := &attributes.ChannelServerListDataContainer{}
	err := Get(l, fmt.Sprintf("http://atlas-nginx:80/ms/wrg/channelServers?world=%d", worldId), r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
