package writer

import (
	"atlas-clc/socket/response"
	"atlas-clc/world"
	"github.com/sirupsen/logrus"
)

const OpCodeServerListRecommendations uint16 = 0x1B

func WriteRecommendedWorlds(l logrus.FieldLogger) func(wrs []world.Recommendation) []byte {
	return func(wrs []world.Recommendation) []byte {
		w := response.NewWriter(l)
		w.WriteShort(OpCodeServerListRecommendations)
		w.WriteByte(byte(len(wrs)))
		for _, x := range wrs {
			w.WriteInt(uint32(x.WorldId()))
			w.WriteAsciiString(x.Reason())
		}
		return w.Bytes()
	}
}
