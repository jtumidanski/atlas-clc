package main

import (
	"atlas-clc/configuration"
	"atlas-clc/logger"
	"atlas-clc/rest"
	"atlas-clc/session"
	"atlas-clc/socket"
	"atlas-clc/tasks"
	"atlas-clc/tracing"
	"context"
	"github.com/opentracing/opentracing-go"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const serviceName = "atlas-clc"

func main() {
	l := logger.CreateLogger()
	l.Infoln("Starting main service.")

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	tc, err := tracing.InitTracer(l)(serviceName)
	if err != nil {
		l.WithError(err).Fatal("Unable to initialize tracer.")
	}
	defer func(tc io.Closer) {
		err := tc.Close()
		if err != nil {
			l.WithError(err).Errorf("Unable to close tracer.")
		}
	}(tc)

	config, err := configuration.GetConfiguration()
	if err != nil {
		l.WithError(err).Fatal("Unable to successfully load configuration.")
	}

	socket.CreateSocketService(l, ctx, wg)

	rest.CreateService(l, ctx, wg, "/ms/clc", session.InitResource)

	go tasks.Register(session.NewTimeout(l, time.Millisecond*time.Duration(config.TimeoutTaskInterval)))

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Infoln("Shutting down via signal:", sig)
	cancel()
	wg.Wait()

	span := opentracing.StartSpan("teardown")
	defer span.Finish()
	session.DestroyAll(l, span, session.GetRegistry())

	l.Infoln("Service shutdown.")
}
