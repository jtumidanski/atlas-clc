package main

import (
   "log"
   "net"
   "os"
)

var sessionId = 0

func main() {
   l := log.New(os.Stdout, "morg ", log.LstdFlags|log.Lmicroseconds)

   l.Println("Starting tcp server on 0.0.0.0:8484")
   lis, err := net.Listen("tcp", "0.0.0.0:8484")
   if err != nil {
      l.Println("Error listening:", err.Error())
      os.Exit(1)
   }
   defer lis.Close()

   for {
      c, err := lis.Accept()
      if err != nil {
         l.Println("Error connecting:", err.Error())
         return
      }
      l.Println("Client connected.")

      go NewConnectionHandler(l).Init(c, sessionId)
      sessionId += 1
   }
}
