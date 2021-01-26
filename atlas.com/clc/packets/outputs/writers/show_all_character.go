package writers

import (
	"atlas-clc/packets/outputs"
	"atlas-clc/packets/outputs/constants"
)

func WriteShowAllCharacter(characters uint32, unk uint32) []byte {
	w := outputs.NewWriter()
	w.WriteShort(constants.ViewAllCharacters)
	if characters > 0 {
		w.WriteByte(1)
	} else {
		w.WriteByte(5)
	}
	w.WriteInt(characters)
	w.WriteInt(unk)
	return w.Bytes()
}
