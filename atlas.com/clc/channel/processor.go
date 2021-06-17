package channel

import (
	"errors"
)

func GetChannels() ([]Model, error) {
	r, err := requestChannels()
	if err != nil {
		return nil, err
	}

	var cs = makeChannelList(r.DataList())
	return cs, nil
}

func GetChannelsForWorld(worldId byte) ([]Model, error) {
	r, err := requestChannelsForWorld(worldId)
	if err != nil {
		return nil, err
	}

	var cs = makeChannelList(r.DataList())
	return cs, nil
}

func GetChannelForWorld(worldId byte, channelId byte) (*Model, error) {
	r, err := requestChannelsForWorld(worldId)
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

func GetChannelLoadByWorld() (map[int][]Load, error) {
	cs, err := GetChannels()
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

func makeChannelList(d []ChannelServerData) []Model {
	var cs = make([]Model, 0)
	for _, x := range d {
		c := makeChannel(x)
		cs = append(cs, c)
	}
	return cs
}

func makeChannel(data ChannelServerData) Model {
	att := data.Attributes
	return NewChannelBuilder().
		SetWorldId(att.WorldId).
		SetChannelId(att.ChannelId).
		SetCapacity(att.Capacity).
		SetIpAddress(att.IpAddress).
		SetPort(att.Port).
		Build()
}
