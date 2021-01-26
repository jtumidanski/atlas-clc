package writers

import (
	"atlas-clc/models"
	"atlas-clc/packets/outputs"
	"atlas-clc/packets/outputs/constants"
)

func WriteRecommendedWorlds(wrs []models.WorldRecommendation) []byte {
	w := outputs.NewWriter()
	w.WriteShort(constants.RecommendedWorldMessage)
	w.WriteByte(byte(len(wrs)))
	for _, x := range wrs {
		w.WriteInt(uint32(x.WorldId()))
		w.WriteAsciiString(x.Reason())
	}
	return w.Bytes()
}
