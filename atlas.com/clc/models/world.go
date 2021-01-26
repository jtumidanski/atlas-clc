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

func NewWorld(id byte, name string, flag int, message string, eventMessage string, recommended bool, recommendedMessage string, capacityStatus uint16) *World {
	return &World{id, name, flag, message, eventMessage, recommended, recommendedMessage, capacityStatus, []ChannelLoad{}}
}

func (w *World) Id() byte {
	return w.id
}

func (w *World) SetChannelLoad(val []ChannelLoad) *World {
	nw := NewWorld(w.id, w.name, w.flag, w.message, w.eventMessage, w.recommended, w.recommendedMessage, w.capacityStatus)
	nw.channelLoad = val
	return nw
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

type WorldRecommendation struct {
	worldId byte
	reason  string
}

func (r WorldRecommendation) WorldId() byte {
	return r.worldId
}

func (r WorldRecommendation) Reason() string {
	return r.reason
}

func NewWorldRecommendation(worldId byte, reason string) *WorldRecommendation {
	return &WorldRecommendation{worldId, reason}
}
