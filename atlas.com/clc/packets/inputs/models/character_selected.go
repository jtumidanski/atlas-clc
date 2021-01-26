package models

type CharacterSelected struct {
	characterId int32
	macs        string
	hwid        string
}

func (s CharacterSelected) CharacterId() int32 {
	return s.characterId
}

func NewCharacterSelected(characterId int32, macs string, hwid string) *CharacterSelected {
	return &CharacterSelected{characterId, macs, hwid}
}
