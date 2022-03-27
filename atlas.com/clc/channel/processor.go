package channel

import (
	"atlas-clc/rest/requests"
	"errors"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func GetAll(l logrus.FieldLogger, span opentracing.Span) ([]Model, error) {
	r, err := requestChannels()(l, span)
	if err != nil {
		return nil, err
	}

	var cs = makeChannelList(r.DataList())
	return cs, nil
}

func GetAllForWorld(l logrus.FieldLogger, span opentracing.Span) func(worldId byte) ([]Model, error) {
	return func(worldId byte) ([]Model, error) {
		r, err := requestChannelsForWorld(worldId)(l, span)
		if err != nil {
			return nil, err
		}

		var cs = makeChannelList(r.DataList())
		return cs, nil
	}
}

func GetForWorldById(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte) (*Model, error) {
	return func(worldId byte, channelId byte) (*Model, error) {
		r, err := requestChannelsForWorld(worldId)(l, span)
		if err != nil {
			return nil, err
		}

		for _, x := range r.DataList() {
			w := makeChannel(x)
			if w.ChannelId() == channelId {
				return &w, nil
			}
		}
		return nil, errors.New("unable to locate channel for world")
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

func makeChannelList(d []requests.DataBody[attributes]) []Model {
	var cs = make([]Model, 0)
	for _, x := range d {
		c := makeChannel(x)
		cs = append(cs, c)
	}
	return cs
}

func makeChannel(data requests.DataBody[attributes]) Model {
	att := data.Attributes
	return NewChannelBuilder().
		SetWorldId(att.WorldId).
		SetChannelId(att.ChannelId).
		SetCapacity(att.Capacity).
		SetIpAddress(att.IpAddress).
		SetPort(att.Port).
		Build()
}
