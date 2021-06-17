package writer

import (
	"atlas-clc/socket/response"
	"github.com/sirupsen/logrus"
)

const OpCodeCharacterNameCheck uint16 = 0x0D

func WriteCharacterNameCheck(l logrus.FieldLogger) func(name string, invalid bool) []byte {
	return func(name string, invalid bool) []byte {
		w := response.NewWriter(l)
		w.WriteShort(OpCodeCharacterNameCheck)
		w.WriteAsciiString(name)
		w.WriteBool(invalid)
		return w.Bytes()
	}
}
