package writer

import (
	"atlas-clc/character"
	"atlas-clc/socket/response"
)

const ViewAllCharacters uint16 = 0x08

func WriteShowAllCharacter(characters uint32, unk uint32) []byte {
	w := response.NewWriter()
	w.WriteShort(ViewAllCharacters)
	if characters > 0 {
		w.WriteByte(1)
	} else {
		w.WriteByte(5)
	}
	w.WriteInt(characters)
	w.WriteInt(unk)
	return w.Bytes()
}

func WriteShowAllCharacterInfo(worldId byte, characters []character.Model, usePIC bool) []byte {
	w := response.NewWriter()
	w.WriteShort(ViewAllCharacters)
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
