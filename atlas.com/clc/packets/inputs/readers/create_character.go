package readers

import (
	"atlas-clc/packets/inputs"
	"atlas-clc/packets/inputs/models"
)

func ReadCreateCharacter(reader *inputs.Reader) *models.CreateCharacter {
	name := reader.ReadAsciiString()
	job := reader.ReadUint32()
	face := reader.ReadUint32()
	hair := reader.ReadUint32()
	hairColor := reader.ReadUint32()
	skinColor := reader.ReadUint32()
	top := reader.ReadUint32()
	bottom := reader.ReadUint32()
	shoes := reader.ReadUint32()
	weapon := reader.ReadUint32()
	gender := reader.ReadByte()
	return models.NewCreateCharacter(name, job, face, hair, hairColor, skinColor, top, bottom, shoes, weapon, gender)
}
