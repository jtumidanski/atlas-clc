package writers

import (
	"atlas-clc/packets/outputs"
	"atlas-clc/packets/outputs/constants"
)

func WritePermanentBan(reason byte) []byte {
	w := outputs.NewWriter()
	w.WriteShort(constants.LoginStatus)
	w.WriteByte(2)
	w.WriteByte(0)
	w.WriteInt(0)
	w.WriteByte(0)
	w.WriteLong(0)
	return w.Bytes()
}
