package channel

import (
	"atlas-clc/rest/requests"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	ServicePrefix   string = "/ms/wrg/"
	Service                = requests.BaseRequest + ServicePrefix
	ChannelResource        = Service + "channelServers/"
	ByWorld                = ChannelResource + "?world=%d"
)

func requestChannels(l logrus.FieldLogger) (*dataContainer, error) {
	r := &dataContainer{}
	err := requests.Get(l)(ChannelResource, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func requestChannelsForWorld(l logrus.FieldLogger) func(worldId byte) (*dataContainer, error) {
	return func(worldId byte) (*dataContainer, error) {
		r := &dataContainer{}
		err := requests.Get(l)(fmt.Sprintf(ByWorld, worldId), r)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}
