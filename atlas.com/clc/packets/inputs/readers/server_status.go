package readers

import (
	"atlas-clc/packets/inputs"
	"atlas-clc/packets/inputs/models"
)

func ReadServerStatus(reader *inputs.Reader) *models.ServerStatus {
	wid := byte(reader.ReadUint16())
	return models.NewServerStatus(wid)
}
