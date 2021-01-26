package models

type ServerStatus struct {
	worldId byte
}

func (s ServerStatus) WorldId() byte {
	return s.worldId
}

func NewServerStatus(worldId byte) *ServerStatus {
	return &ServerStatus{worldId}
}
