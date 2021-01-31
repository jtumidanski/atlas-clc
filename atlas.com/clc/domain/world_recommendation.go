package domain

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

func NewWorldRecommendation(worldId byte, reason string) WorldRecommendation {
	return WorldRecommendation{worldId, reason}
}