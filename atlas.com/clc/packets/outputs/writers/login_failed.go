package writers

import (
	"atlas-clc/packets/outputs"
	"atlas-clc/packets/outputs/constants"
)

func WriteLoginFailed(reason byte) []byte {
	w := outputs.NewWriter()
	w.WriteShort(constants.LoginStatus)
	w.WriteByte(reason)
	w.WriteByte(0)
	w.WriteInt(0)
	return w.Bytes()
}
