package requests

import (
	"atlas-clc/rest/attributes"
	"fmt"
	"log"
)

func GetChannels(l *log.Logger) (*attributes.ChannelServerDataContainer, error) {
	r := &attributes.ChannelServerDataContainer{}
	err := Get(l, "http://atlas-nginx:80/ms/wrg/channelServers/", r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetChannelsForWorld(l *log.Logger, worldId byte) (*attributes.ChannelServerDataContainer, error) {
	r := &attributes.ChannelServerDataContainer{}
	err := Get(l, fmt.Sprintf("http://atlas-nginx:80/ms/wrg/channelServers/?world=%d", worldId), r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
