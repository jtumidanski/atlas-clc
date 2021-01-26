package writers

import (
	"atlas-clc/models"
	"atlas-clc/packets/outputs"
	"atlas-clc/packets/outputs/constants"
)

func WriteShowAllCharacterInfo(worldId byte, characters []models.Character, usePIC bool) []byte {
	w := outputs.NewWriter()
	w.WriteShort(constants.ViewAllCharacters)
	w.WriteByte(0)
	w.WriteByte(worldId)
	w.WriteByte(byte(len(characters)))
	for _, x := range characters {
		WriteCharacter(w, x, true)
	}
	if usePIC {
		w.WriteByte(1)
	} else {
		w.WriteByte(2)
	}
	return w.Bytes()
}
