package writers

import (
	"atlas-clc/packets/outputs"
	"atlas-clc/packets/outputs/constants"
	"strconv"
	"strings"
)

func WriteServerIp(ipAddress string, port uint16, characterId uint32) []byte {
	w := outputs.NewWriter()
	w.WriteShort(constants.ServerIp)
	w.WriteShort(0)
	ob := ipAsByteArray(ipAddress)
	w.WriteByteArray(ob)
	w.WriteShort(port)
	w.WriteInt(characterId)
	w.WriteByteArray([]byte{0, 0, 0, 0, 0})
	return w.Bytes()
}

func ipAsByteArray(ipAddress string) []byte {
	var ob = make([]byte, 0)
	os := strings.Split(ipAddress, ".")
	for _, x := range os {
		o, err := strconv.ParseUint(x, 10, 8)
		if err == nil {
			ob = append(ob, byte(o))
		}
	}
	return ob
}
