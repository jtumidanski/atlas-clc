package writers

import (
	"atlas-clc/packets/outputs"
	"atlas-clc/packets/outputs/constants"
)

func WriteCharacterNameCheckResponse(name string, invalid bool) []byte {
	w := outputs.NewWriter()
	w.WriteShort(constants.CharacterNameResponse)
	w.WriteAsciiString(name)
	w.WriteBool(invalid)
	return w.Bytes()
}
