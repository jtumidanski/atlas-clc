package models

type CharacterListRequest struct {
	worldId   byte
	channelId byte
}

func (r CharacterListRequest) WorldId() byte {
	return r.worldId
}

func (r CharacterListRequest) ChannelId() byte {
	return r.channelId
}

func NewCharacterListRequest(worldId byte, channelId byte) *CharacterListRequest {
	return &CharacterListRequest{worldId, channelId}
}
