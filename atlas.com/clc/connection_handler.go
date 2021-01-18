package main

import (
   "atlas-clc/models"
   "atlas-clc/registries"
   "bufio"
   "log"
   "net"
)

type ConnectionHandler struct {
   l *log.Logger
}

func NewConnectionHandler(l *log.Logger) *ConnectionHandler {
   return &ConnectionHandler{l}
}

func (h *ConnectionHandler) Init(c net.Conn, sessionId int) {
   h.l.Println("Client " + c.RemoteAddr().String() + " connected.")

   s := models.NewSession(sessionId, &c, h.l)
   registries.GetSessionRegistry().AddSession(s)
   s.WriteHello()

   h.ReadLoop(c, sessionId)
}

func (h *ConnectionHandler) ReadLoop(c net.Conn, sessionId int) {
   bufio.NewReader(c).Size()
   bufio.NewReader(c).Peek()
   buffer, err := bufio.NewReader(c).ReadBytes('\n')
   if err != nil {
      h.l.Println("Client left.")
      c.Close()
      return
   }
   h.l.Println("Client message:", string(buffer[:len(buffer)-1]))
   h.ReadLoop(c, sessionId)
}
