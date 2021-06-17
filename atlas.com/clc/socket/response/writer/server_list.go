package writer

import (
	"atlas-clc/channel"
	"atlas-clc/socket/response"
	"fmt"
	"github.com/sirupsen/logrus"
)

const OpCodeServerList uint16 = 0x0A

func WriteServerListEntry(l logrus.FieldLogger) func(worldId byte, worldName string, flag int, eventMessage string, channelLoad []channel.Load) []byte {
	return func(worldId byte, worldName string, flag int, eventMessage string, channelLoad []channel.Load) []byte {
		w := response.NewWriter(l)
		w.WriteShort(OpCodeServerList)
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
}

func WriteServerListEnd(l logrus.FieldLogger) []byte {
	w := response.NewWriter(l)
	w.WriteShort(OpCodeServerList)
	w.WriteByte(byte(0xFF))
	return w.Bytes()
}
