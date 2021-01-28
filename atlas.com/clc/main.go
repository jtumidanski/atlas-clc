package main

import (
	"atlas-clc/registries"
	rest2 "atlas-clc/rest"
	sessions2 "atlas-clc/sessions"
	"atlas-clc/socket"
	"atlas-clc/tasks"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	l := log.New(os.Stdout, "clc ", log.LstdFlags|log.Lmicroseconds)

	config, err := registries.GetConfiguration()
	if err != nil {
		l.Fatal("[ERROR] Unable to successfully load configuration.")
	}

	ss, err := socket.NewServer(l, sessions2.NewSessionService(l), socket.IpAddress("0.0.0.0"), socket.Port(8484))
	if err != nil {
		return
	}
	go ss.Run()

	hs := rest2.NewServer(l)
	go hs.Run()

	go tasks.Register(tasks.NewTimeout(l, time.Second*time.Duration(config.TimeoutTaskInterval)))

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Println("[INFO] Shutting down via signal:", sig)

	sessions := registries.GetSessionRegistry().GetAll()
	for _, s := range sessions {
		socket.Disconnect(l, &s)
		s.Disconnect()
	}
}
