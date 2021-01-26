package models

type ViewAllCharactersSelected struct {
	characterId int32
	worldId     int32
	macs        string
	hwid        string
}

func (s ViewAllCharactersSelected) CharacterId() int32 {
	return s.characterId
}

func (s ViewAllCharactersSelected) WorldId() int32 {
	return s.worldId
}

func NewViewAllCharactersSelected(characterId int32, worldId int32, macs string, hwid string) *ViewAllCharactersSelected {
	return &ViewAllCharactersSelected{characterId, worldId, macs, hwid}
}
