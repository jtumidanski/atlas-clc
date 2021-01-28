package writer

import (
	"atlas-clc/domain"
	"atlas-clc/socket/response"
)

const OpCodeServerListRecommendations uint16 = 0x1B

func WriteRecommendedWorlds(wrs []domain.WorldRecommendation) []byte {
	w := response.NewWriter()
	w.WriteShort(OpCodeServerListRecommendations)
	w.WriteByte(byte(len(wrs)))
	for _, x := range wrs {
		w.WriteInt(uint32(x.WorldId()))
		w.WriteAsciiString(x.Reason())
	}
	return w.Bytes()
}
