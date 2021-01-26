package writers

import (
	"atlas-clc/packets/outputs"
	"atlas-clc/packets/outputs/constants"
)

func WriteServerIp(ipAddress string, port uint16, characterId uint32) []byte {
	w := outputs.NewWriter()
	w.WriteShort(constants.ServerIp)
	w.WriteShort(0)
	w.WriteByteArray([]byte(ipAddress))
	w.WriteShort(port)
	w.WriteInt(characterId)
	w.WriteByteArray([]byte{0, 0, 0, 0, 0})
	return w.Bytes()
}
