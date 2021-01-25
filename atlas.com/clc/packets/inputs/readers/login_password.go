package readers

import (
   "atlas-clc/packets"
   "atlas-clc/packets/inputs/models"
)

func ReadLoginPassword(reader *packets.Reader) *models.LoginPassword {
   login := reader.ReadAsciiString()
   password := reader.ReadAsciiString()
   reader.Skip(6)
   hwid := reader.ReadBytes(4)
   return models.NewLoginPassword(login, password, hwid)
}
