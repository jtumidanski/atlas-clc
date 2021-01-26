package writers

import (
	"atlas-clc/models"
	"atlas-clc/packets/outputs"
	"atlas-clc/packets/outputs/constants"
)

func WriteNewCharacter(character models.Character) []byte {
	w := outputs.NewWriter()
	w.WriteShort(constants.AddNewCharacterEntry)
	w.WriteByte(0)
	WriteCharacter(w, character, false)
	return w.Bytes()
}
