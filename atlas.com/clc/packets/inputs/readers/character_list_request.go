package readers

import (
	"atlas-clc/packets/inputs"
	"atlas-clc/packets/inputs/models"
)

func ReadCharacterListRequest(reader *inputs.Reader) *models.CharacterListRequest {
	reader.ReadByte()
	return models.NewCharacterListRequest(reader.ReadByte(), reader.ReadByte() + 1)
}
