package writer

import "atlas-clc/socket/response"

const OpCodeServerStatus uint16 = 0x03

func WriteWorldCapacityStatus(status uint16) []byte {
	w := response.NewWriter()
	w.WriteShort(OpCodeServerStatus)
	w.WriteShort(status)
	return w.Bytes()
}