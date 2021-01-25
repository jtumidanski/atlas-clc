package writers

import (
	"atlas-clc/packets/outputs"
	"atlas-clc/packets/outputs/constants"
)

func WriteTemporaryBan(until uint64, reason byte) []byte {
	w := outputs.NewWriter()
	w.WriteShort(constants.LoginStatus)
	w.WriteByte(2)
	w.WriteByte(0)
	w.WriteInt(0)
	w.WriteByte(reason)
	// Temp ban date is handled as a 64-bit long, number of 100NS intervals since 1/1/1601.
	w.WriteLong(until)
	return w.Bytes()
}
