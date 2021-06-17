package writer

import (
	"atlas-clc/socket/response"
	"github.com/sirupsen/logrus"
)

const OpCodeServerStatus uint16 = 0x03

func WriteWorldCapacityStatus(l logrus.FieldLogger) func(status uint16) []byte {
	return func(status uint16) []byte {
		w := response.NewWriter(l)
		w.WriteShort(OpCodeServerStatus)
		w.WriteShort(status)
		return w.Bytes()
	}
}
