package handlers

import (
   "atlas-clc/packets"
   "atlas-clc/packets/inputs/constants"
   "log"
)

func GetHandler(op int16) func(sessionId int, r *packets.Reader, l *log.Logger) {
   switch op {
   case constants.LoginPassword:
      return HandleLoginPassword
   }
   return nil
}
