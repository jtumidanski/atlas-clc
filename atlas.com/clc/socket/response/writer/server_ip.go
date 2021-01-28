package writer

import (
	"atlas-clc/socket/response"
	"strconv"
	"strings"
)

const OpCodeServerIp uint16 = 0x0C

func WriteServerIp(ipAddress string, port uint16, characterId uint32) []byte {
	w := response.NewWriter()
	w.WriteShort(OpCodeServerIp)
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
