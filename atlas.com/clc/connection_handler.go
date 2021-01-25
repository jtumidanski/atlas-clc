package main

import (
   "atlas-clc/crypto"
   "atlas-clc/packets"
   "atlas-clc/packets/inputs/handlers"
   "atlas-clc/registries"
   "atlas-clc/sessions"
   "log"
   "net"
   "time"
)

type ConnectionHandler struct {
   l *log.Logger
}

func NewConnectionHandler(l *log.Logger) *ConnectionHandler {
   return &ConnectionHandler{l}
}

func (ch *ConnectionHandler) Init(c net.Conn, sessionId int) {
   ch.l.Println("Client " + c.RemoteAddr().String() + " connected.")

   s := sessions.NewSession(sessionId, &c, ch.l)
   registries.GetSessionRegistry().AddSession(s)
   s.WriteHello()

   ch.ReadLoop(c, sessionId, 4)
}

func (ch *ConnectionHandler) ReadLoop(c net.Conn, sessionId int, headerSize int) {
   header := true
   readSize := headerSize

   session := registries.GetSessionRegistry().GetSession(sessionId)

   for {
      buffer := make([]byte, readSize)

      if _, err := c.Read(buffer); err != nil {
         break
      }

      if header {
         readSize = crypto.GetPacketLength(buffer)
      } else {
         readSize = headerSize

         result := buffer
         if session.GetRecv() != nil {
            ue := session.GetRecv().Decrypt(buffer, true, true)
            result = ue
         }
         ch.Handle(sessionId, result)
      }

      header = !header
   }

   ch.l.Printf("Session %d exiting read loop.", sessionId)
}

func (ch *ConnectionHandler) Handle(sessionId int, p packets.Packet) {
   go func(sessionId int, reader packets.Reader) {
      op := reader.ReadInt16()
      h := handlers.GetHandler(op)
      if h != nil {
         h(sessionId, &reader, ch.l)
      } else {
         ch.l.Printf("Session %d read a message with op %d.", sessionId, op)
      }
   }(sessionId, packets.NewReader(&p, time.Now().Unix()))
}
