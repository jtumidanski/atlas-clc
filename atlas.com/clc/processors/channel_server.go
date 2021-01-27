package processors

import (
	"atlas-clc/models"
	"atlas-clc/rest/attributes"
	"atlas-clc/rest/requests"
	"errors"
	"log"
)

func GetChannels(l *log.Logger) ([]models.Channel, error) {
	r, err := requests.GetChannels(l)
	if err != nil {
		return nil, err
	}

	var cs = makeChannelList(r.DataList())
	return cs, nil
}

func GetChannelsForWorld(l *log.Logger, worldId byte) ([]models.Channel, error) {
	r, err := requests.GetChannelsForWorld(l, worldId)
	if err != nil {
		return nil, err
	}

	var cs = makeChannelList(r.DataList())
	return cs, nil
}

func GetChannelForWorld(l *log.Logger, worldId byte, channelId byte) (*models.Channel, error) {
	r, err := requests.GetChannelsForWorld(l, worldId)
	if err != nil {
		return nil, err
	}

	for _, x := range r.DataList() {
		w := makeChannel(x)
		if w.ChannelId() == channelId {
			return w, nil
		}
	}
	return nil, errors.New("unable to locate channel for world")
}

func GetChannelLoadByWorld(l *log.Logger) (map[int][]models.ChannelLoad, error) {
	cs, err := GetChannels(l)
	if err != nil {
		return nil, err
	}

	var cls = make(map[int][]models.ChannelLoad, 0)
	for _, x := range cs {
		cl := models.NewChannelLoad(x.ChannelId(), x.Capacity())
		if _, ok := cls[int(x.WorldId())]; ok {
			cls[int(x.WorldId())] = append(cls[int(x.WorldId())], *cl)
		} else {
			cls[int(x.WorldId())] = append([]models.ChannelLoad{}, *cl)
		}
	}
	return cls, nil
}

func makeChannelList(d []attributes.ChannelServerData) []models.Channel {
	var cs = make([]models.Channel, 0)
	for _, x := range d {
		c := makeChannel(x)
		cs = append(cs, *c)
	}
	return cs
}

func makeChannel(data attributes.ChannelServerData) *models.Channel {
	att := data.Attributes
	return models.NewChannelBuilder().
		SetWorldId(att.WorldId).
		SetChannelId(att.ChannelId).
		SetCapacity(att.Capacity).
		SetIpAddress(att.IpAddress).
		SetPort(att.Port).
		Build()
}
