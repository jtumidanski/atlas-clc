package channel

import (
	"atlas-clc/model"
	"atlas-clc/rest/requests"
	"errors"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func GetAll(l logrus.FieldLogger, span opentracing.Span) ([]Model, error) {
	return requests.SliceProvider[attributes, Model](l, span)(requestChannels(), makeChannel)()
}

func ByWorldModelProvider(l logrus.FieldLogger, span opentracing.Span) func(worldId byte) model.SliceProvider[Model] {
	return func(worldId byte) model.SliceProvider[Model] {
		return requests.SliceProvider[attributes, Model](l, span)(requestChannelsForWorld(worldId), makeChannel)
	}
}

func GetAllForWorld(l logrus.FieldLogger, span opentracing.Span) func(worldId byte) ([]Model, error) {
	return func(worldId byte) ([]Model, error) {
		return ByWorldModelProvider(l, span)(worldId)()
	}
}

func GetRandomChannelForWorld(l logrus.FieldLogger, span opentracing.Span) func(worldId byte) (Model, error) {
	return func(worldId byte) (Model, error) {
		return model.SliceProviderToProviderAdapter(ByWorldModelProvider(l, span)(worldId), randomChannelFilter)()
	}
}

func randomChannelFilter(ms []Model) (Model, error) {
	rand.Seed(time.Now().Unix())
	return ms[rand.Intn(len(ms))], nil
}

func GetForWorldById(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte) (Model, error) {
	return func(worldId byte, channelId byte) (Model, error) {
		return model.SliceProviderToProviderAdapter(ByWorldModelProvider(l, span)(worldId), matchingChannelFilter(channelId))()
	}
}

func matchingChannelFilter(channelId byte) model.PreciselyOneFilter[Model] {
	return func(ms []Model) (Model, error) {
		for _, m := range ms {
			if m.ChannelId() == channelId {
				return m, nil
			}
		}
		return Model{}, errors.New("unable to locate channel for world")
	}
}

func GetChannelLoadByWorld(l logrus.FieldLogger, span opentracing.Span) (map[int][]Load, error) {
	cs, err := GetAll(l, span)
	if err != nil {
		return nil, err
	}

	var cls = make(map[int][]Load, 0)
	for _, x := range cs {
		cl := NewChannelLoad(x.ChannelId(), x.Capacity())
		if _, ok := cls[int(x.WorldId())]; ok {
			cls[int(x.WorldId())] = append(cls[int(x.WorldId())], cl)
		} else {
			cls[int(x.WorldId())] = append([]Load{}, cl)
		}
	}
	return cls, nil
}

func makeChannel(data requests.DataBody[attributes]) (Model, error) {
	att := data.Attributes
	return NewChannelBuilder().
		SetWorldId(att.WorldId).
		SetChannelId(att.ChannelId).
		SetCapacity(att.Capacity).
		SetIpAddress(att.IpAddress).
		SetPort(att.Port).
		Build(), nil
}
