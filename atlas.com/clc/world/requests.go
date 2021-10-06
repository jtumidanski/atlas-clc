package world

import (
	"atlas-clc/rest/requests"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	ServicePrefix  string = "/ms/wrg/"
	Service               = requests.BaseRequest + ServicePrefix
	WorldsResource        = Service + "worlds/"
	WorldsById            = WorldsResource + "%d"
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

func requestWorlds() Request {
	return makeRequest(WorldsResource)
}

func requestWorld(worldId byte) Request {
	return makeRequest(fmt.Sprintf(WorldsById, worldId))
}
