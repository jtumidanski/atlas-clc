package writers

import (
   "atlas-clc/packets/outputs"
)

func WriteHello(version uint16, sendIv []byte, recvIv []byte) []byte {
   w := outputs.NewWriter()
   w.WriteShort(uint16(0x0E))
   w.WriteShort(version)
   w.WriteShort(1)
   w.WriteByte(49)
   w.WriteByteArray(recvIv)
   w.WriteByteArray(sendIv)
   w.WriteByte(8)
   return w.Bytes()
}
