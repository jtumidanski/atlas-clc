package writer

import (
	"atlas-clc/socket/response"
)

const OpCodeCharacterNameCheck uint16 = 0x0D

func WriteCharacterNameCheck(name string, invalid bool) []byte {
	w := response.NewWriter()
	w.WriteShort(OpCodeCharacterNameCheck)
	w.WriteAsciiString(name)
	w.WriteBool(invalid)
	return w.Bytes()
}
