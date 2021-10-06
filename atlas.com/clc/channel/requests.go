package channel

import (
	"atlas-clc/rest/requests"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	ServicePrefix   string = "/ms/wrg/"
	Service                = requests.BaseRequest + ServicePrefix
	ChannelResource        = Service + "channelServers/"
	ByWorld                = ChannelResource + "?world=%d"
)

type Request func(l logrus.FieldLogger, span opentracing.Span) (*dataContainer, error)

func makeRequest(url string) Request {
	return func(l logrus.FieldLogger, span opentracing.Span) (*dataContainer, error) {
		ar := &dataContainer{}
		err := requests.Get(l, span)(url, ar)
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
