package models

type World struct {
	id                 byte
	name               string
	flag               int
	message            string
	eventMessage       string
	recommended        bool
	recommendedMessage string
	capacityStatus     uint16
	channelLoad        []ChannelLoad
}

func (w *World) Id() byte {
	return w.id
}

func (w *World) SetChannelLoad(val []ChannelLoad) *World {
	return CloneWorld(w).
		SetChannelLoad(val).
		Build()
}

func (w *World) Name() string {
	return w.name
}

func (w *World) Flag() int {
	return w.flag
}

func (w *World) EventMessage() string {
	return w.eventMessage
}

func (w *World) ChannelLoad() []ChannelLoad {
	return w.channelLoad
}

func (w *World) Recommended() bool {
	return w.recommended
}

func (w *World) Recommendation() *WorldRecommendation {
	return NewWorldRecommendation(w.id, w.recommendedMessage)
}

func (w *World) CapacityStatus() uint16 {
	return w.capacityStatus
}

type worldBuilder struct {
	id                 byte
	name               string
	flag               int
	message            string
	eventMessage       string
	recommended        bool
	recommendedMessage string
	capacityStatus     uint16
	channelLoad        []ChannelLoad
}

func NewWorldBuilder() *worldBuilder {
	return &worldBuilder{}
}

func CloneWorld(o *World) *worldBuilder {
	return &worldBuilder{
		id:                 o.id,
		name:               o.name,
		flag:               o.flag,
		message:            o.message,
		eventMessage:       o.eventMessage,
		recommended:        o.recommended,
		recommendedMessage: o.recommendedMessage,
		capacityStatus:     o.capacityStatus,
		channelLoad:        o.channelLoad,
	}
}

func (w *worldBuilder) SetId(id byte) *worldBuilder {
	w.id = id
	return w
}

func (w *worldBuilder) SetName(name string) *worldBuilder {
	w.name = name
	return w
}

func (w *worldBuilder) SetFlag(flag int) *worldBuilder {
	w.flag = flag
	return w
}

func (w *worldBuilder) SetMessage(message string) *worldBuilder {
	w.message = message
	return w
}

func (w *worldBuilder) SetEventMessage(eventMessage string) *worldBuilder {
	w.eventMessage = eventMessage
	return w
}

func (w *worldBuilder) SetRecommended(recommended bool) *worldBuilder {
	w.recommended = recommended
	return w
}

func (w *worldBuilder) SetRecommendedMessage(recommendedMessage string) *worldBuilder {
	w.recommendedMessage = recommendedMessage
	return w
}

func (w *worldBuilder) SetCapacityStatus(capacityStatus uint16) *worldBuilder {
	w.capacityStatus = capacityStatus
	return w
}

func (w *worldBuilder) SetChannelLoad(channelLoad []ChannelLoad) *worldBuilder {
	w.channelLoad = channelLoad
	return w
}

func (w *worldBuilder) Build() *World {
	return &World{
		id:                 w.id,
		name:               w.name,
		flag:               w.flag,
		message:            w.message,
		eventMessage:       w.eventMessage,
		recommended:        w.recommended,
		recommendedMessage: w.recommendedMessage,
		capacityStatus:     w.capacityStatus,
		channelLoad:        w.channelLoad,
	}
}
