package writer

import (
	"atlas-clc/socket/response"
	"github.com/sirupsen/logrus"
)

const OpCodeServerLastConnected uint16 = 0x1A

func WriteSelectWorld(l logrus.FieldLogger) func(worldId int) []byte {
	return func(worldId int) []byte {
		//According to GMS, it should be the world that contains the most characters (most active)
		w := response.NewWriter(l)
		w.WriteShort(OpCodeServerLastConnected)
		w.WriteInt(uint32(worldId))
		return w.Bytes()
	}
}
