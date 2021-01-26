package writers

import (
	"atlas-clc/models"
	"atlas-clc/packets/outputs"
	"atlas-clc/packets/outputs/constants"
	"fmt"
)

func WriteServerListEntry(worldId byte, worldName string, flag int, eventMessage string, channelLoad []models.ChannelLoad) []byte {
	w := outputs.NewWriter()
	w.WriteShort(constants.ServerList)
	w.WriteByte(worldId)
	w.WriteAsciiString(worldName)
	w.WriteByte(byte(flag))
	w.WriteAsciiString(eventMessage)
	w.WriteByte(100) // rate modifier, don't ask O.O!
	w.WriteByte(0)   // event xp * 2.6 O.O!
	w.WriteByte(100) // rate modifier, don't ask O.O!
	w.WriteByte(0)   // drop rate * 2.6
	w.WriteByte(0)
	w.WriteByte(byte(len(channelLoad)))
	for _, x := range channelLoad {
		w.WriteAsciiString(fmt.Sprintf("%s - %d", worldName, x.ChannelId()))
		w.WriteInt(uint32(x.Capacity()))
		w.WriteByte(1)
		w.WriteByte(x.ChannelId() - 1)
		w.WriteBool(false) // adult channel
	}
	w.WriteShort(0)
	return w.Bytes()
}
