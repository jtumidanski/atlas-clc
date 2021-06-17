package writer

import (
	"atlas-clc/character"
	"atlas-clc/socket/response"
	"github.com/sirupsen/logrus"
)

const OpCodeCharacterViewAddNew uint16 = 0x0E

func WriteCharacterViewAddNew(l logrus.FieldLogger) func(character character.Model) []byte {
	return func(character character.Model) []byte {
		w := response.NewWriter(l)
		w.WriteShort(OpCodeCharacterViewAddNew)
		w.WriteByte(0)
		WriteCharacter(w, character, false)
		return w.Bytes()
	}
}
