package processors

import (
	"atlas-clc/domain"
	"atlas-clc/rest/attributes"
	"atlas-clc/rest/requests"
	"errors"
)

func GetChannels() ([]domain.Channel, error) {
	r, err := requests.GetChannels()
	if err != nil {
		return nil, err
	}

	var cs = makeChannelList(r.DataList())
	return cs, nil
}

func GetChannelsForWorld(worldId byte) ([]domain.Channel, error) {
	r, err := requests.GetChannelsForWorld(worldId)
	if err != nil {
		return nil, err
	}

	var cs = makeChannelList(r.DataList())
	return cs, nil
}

func GetChannelForWorld(worldId byte, channelId byte) (*domain.Channel, error) {
	r, err := requests.GetChannelsForWorld(worldId)
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

func GetChannelLoadByWorld() (map[int][]domain.ChannelLoad, error) {
	cs, err := GetChannels()
	if err != nil {
		return nil, err
	}

	var cls = make(map[int][]domain.ChannelLoad, 0)
	for _, x := range cs {
		cl := domain.NewChannelLoad(x.ChannelId(), x.Capacity())
		if _, ok := cls[int(x.WorldId())]; ok {
			cls[int(x.WorldId())] = append(cls[int(x.WorldId())], cl)
		} else {
			cls[int(x.WorldId())] = append([]domain.ChannelLoad{}, cl)
		}
	}
	return cls, nil
}

func makeChannelList(d []attributes.ChannelServerData) []domain.Channel {
	var cs = make([]domain.Channel, 0)
	for _, x := range d {
		c := makeChannel(x)
		cs = append(cs, c)
	}
	return cs
}

func makeChannel(data attributes.ChannelServerData) domain.Channel {
	att := data.Attributes
	return domain.NewChannelBuilder().
		SetWorldId(att.WorldId).
		SetChannelId(att.ChannelId).
		SetCapacity(att.Capacity).
		SetIpAddress(att.IpAddress).
		SetPort(att.Port).
		Build()
}
