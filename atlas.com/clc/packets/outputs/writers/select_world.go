package writers

import (
	"atlas-clc/packets/outputs"
	"atlas-clc/packets/outputs/constants"
)

func WriteSelectWorld(worldId int) []byte {
	//According to GMS, it should be the world that contains the most characters (most active)
	w := outputs.NewWriter()
	w.WriteShort(constants.LastConnectedWorld)
	w.WriteInt(uint32(worldId))
	return w.Bytes()
}
