package readers

import (
	"atlas-clc/packets/inputs"
	"atlas-clc/packets/inputs/models"
)

func ReadCharacterSelected(reader *inputs.Reader) *models.CharacterSelected {
	cid := reader.ReadInt32()
	macs := reader.ReadAsciiString()
	hwid := reader.ReadAsciiString()

	return models.NewCharacterSelected(cid, macs, hwid)
}
