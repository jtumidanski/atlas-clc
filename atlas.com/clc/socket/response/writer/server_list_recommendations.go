package writer

import (
	"atlas-clc/socket/response"
	"atlas-clc/world"
)

const OpCodeServerListRecommendations uint16 = 0x1B

func WriteRecommendedWorlds(wrs []world.Recommendation) []byte {
	w := response.NewWriter()
	w.WriteShort(OpCodeServerListRecommendations)
	w.WriteByte(byte(len(wrs)))
	for _, x := range wrs {
		w.WriteInt(uint32(x.WorldId()))
		w.WriteAsciiString(x.Reason())
	}
	return w.Bytes()
}
