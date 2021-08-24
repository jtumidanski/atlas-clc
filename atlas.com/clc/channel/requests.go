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

type Request func(l logrus.FieldLogger) (*dataContainer, error)

func makeRequest(url string) Request {
	return func(l logrus.FieldLogger) (*dataContainer, error) {
		ar := &dataContainer{}
		err := requests.Get(l)(url, ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}

func requestChannels() Request {
	return makeRequest(ChannelResource)
}

func requestChannelsForWorld(worldId byte) Request {
	return makeRequest(fmt.Sprintf(ByWorld, worldId))
}
