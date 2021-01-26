package socket

import (
   "log"
   "net"
   "os"
)

type Server struct {
   l         *log.Logger
   sessionId int
}

func NewServer(l *log.Logger) *Server {
   return &Server{l, 0}
}

func (s *Server) Run() {
   s.l.Println("Starting tcp server on 0.0.0.0:8484")
   lis, err := net.Listen("tcp", "0.0.0.0:8484")
   if err != nil {
      s.l.Println("Error listening:", err.Error())
      os.Exit(1)
   }
   defer lis.Close()

   for {
      c, err := lis.Accept()
      if err != nil {
         s.l.Println("Error connecting:", err.Error())
         return
      }
      s.l.Println("Client connected.")

      go NewConnectionHandler(s.l).Init(c, s.sessionId)
      s.sessionId += 1
   }
}
