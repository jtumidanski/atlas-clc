package writer

import (
	"atlas-clc/character"
	"atlas-clc/socket/response"
)

const OpCodeCharacterViewAddNew uint16 = 0x0E

func WriteCharacterViewAddNew(character character.Model) []byte {
	w := response.NewWriter()
	w.WriteShort(OpCodeCharacterViewAddNew)
	w.WriteByte(0)
	WriteCharacter(w, character, false)
	return w.Bytes()
}
