package main

import (
	"atlas-clc/configuration"
	"atlas-clc/logger"
	"atlas-clc/rest"
	"atlas-clc/session"
	"atlas-clc/socket"
	"atlas-clc/tasks"
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	l := logger.CreateLogger()
	l.Infoln("Starting main service.")

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	config, err := configuration.GetConfiguration()
	if err != nil {
		l.WithError(err).Fatal("Unable to successfully load configuration.")
	}

	lss := session.NewMapleSessionService(l)

	socket.CreateSocketService(l, lss, ctx, wg)

	rest.CreateRestService(l, ctx, wg)

	go tasks.Register(session.NewTimeout(l, lss, time.Second*time.Duration(config.TimeoutTaskInterval)))

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Infoln("Shutting down via signal:", sig)
	cancel()
	wg.Wait()

	sessions := session.GetRegistry().GetAll()
	for _, s := range sessions {
		lss.Destroy(s.SessionId())
	}

	l.Infoln("Service shutdown.")
}
