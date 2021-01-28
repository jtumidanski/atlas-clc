package writer

import (
	"atlas-clc/socket/response"
)

const OpCodeCharacterNameCheckResponse uint16 = 0x0D

func WriteCharacterNameCheckResponse(name string, invalid bool) []byte {
	w := response.NewWriter()
	w.WriteShort(OpCodeCharacterNameCheckResponse)
	w.WriteAsciiString(name)
	w.WriteBool(invalid)
	return w.Bytes()
}
