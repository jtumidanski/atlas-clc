package writers

import (
	"atlas-clc/packets/outputs"
	"atlas-clc/packets/outputs/constants"
)

func WriteWorldCapacityStatus(status uint16) []byte {
	w := outputs.NewWriter()
	w.WriteShort(constants.ServerStatus)
	w.WriteShort(status)
	return w.Bytes()
}
