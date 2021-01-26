package main

import (
   "atlas-clc/servers/rest"
   "atlas-clc/servers/socket"
   "log"
   "os"
   "os/signal"
   "syscall"
)


func main() {
   l := log.New(os.Stdout, "clc ", log.LstdFlags|log.Lmicroseconds)

   ss := socket.NewServer(l)
   go ss.Run()

   hs := rest.NewServer(l)
   go hs.Run()

   // trap sigterm or interrupt and gracefully shutdown the server
   c := make(chan os.Signal, 1)
   signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

   // Block until a signal is received.
   sig := <-c
   l.Println("Got signal:", sig)
}
