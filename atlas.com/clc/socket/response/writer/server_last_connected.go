package writer

import (
	"atlas-clc/socket/response"
)

const OpCodeServerLastConnected uint16 = 0x1A

func WriteSelectWorld(worldId int) []byte {
	//According to GMS, it should be the world that contains the most characters (most active)
	w := response.NewWriter()
	w.WriteShort(OpCodeServerLastConnected)
	w.WriteInt(uint32(worldId))
	return w.Bytes()
}
