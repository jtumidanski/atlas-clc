package writers

import (
	"atlas-clc/packets/outputs"
	"atlas-clc/packets/outputs/constants"
)

func WriteServerListEnd() []byte {
	w := outputs.NewWriter()
	w.WriteShort(constants.ServerList)
	w.WriteByte(byte(0xFF))
	return w.Bytes()
}
