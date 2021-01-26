package readers

import (
	"atlas-clc/packets/inputs"
	"atlas-clc/packets/inputs/models"
)

func ReadCheckCharacterName(reader *inputs.Reader) *models.CheckCharacterName {
	return models.NewCheckCharacterName(reader.ReadAsciiString())
}
