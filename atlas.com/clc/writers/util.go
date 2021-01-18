package writers

import (
   "bytes"
   "encoding/binary"
   "log"
)

func WriteShort(buf *bytes.Buffer, val uint16) {
   err := binary.Write(buf, binary.LittleEndian, val)
   if err != nil {
      log.Fatal("[ERROR] writing short value")
   }
}

func WriteByte(buf *bytes.Buffer, val byte) {
   err := binary.Write(buf, binary.LittleEndian, val)
   if err != nil {
      log.Fatal("[ERROR] writing byte value")
   }
}

func WriteByteArray(buf *bytes.Buffer, vals []byte) {
   for i := 0; i < len(vals); i++ {
      err := binary.Write(buf, binary.LittleEndian, vals[i])
      if err != nil {
         log.Fatal("[ERROR] writing byte value")
      }
   }
}
