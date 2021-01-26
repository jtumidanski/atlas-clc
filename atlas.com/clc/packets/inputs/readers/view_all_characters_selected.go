package readers

import (
	"atlas-clc/packets/inputs"
	"atlas-clc/packets/inputs/models"
)

func ReadViewAllCharactersSelected(reader *inputs.Reader) *models.ViewAllCharactersSelected {
	cid := reader.ReadInt32()
	wid := reader.ReadInt32()
	macs := reader.ReadAsciiString()
	hwid := reader.ReadAsciiString()
	return models.NewViewAllCharactersSelected(cid, wid, macs, hwid)
}
