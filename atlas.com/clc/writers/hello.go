package writers

import (
   "bytes"
)

func WriteHello(version uint16, sendIv []byte, recvIv []byte) *bytes.Buffer {
   buf := new(bytes.Buffer)
   WriteShort(buf, uint16(0x0E))
   WriteShort(buf, version)
   WriteShort(buf, 1)
   WriteByte(buf, 49)
   WriteByteArray(buf, recvIv)
   WriteByteArray(buf, sendIv)
   WriteByte(buf, 8)
   return buf
}
