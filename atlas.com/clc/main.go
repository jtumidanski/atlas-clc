package main

import (
	"atlas-clc/configuration"
	"atlas-clc/logger"
	"atlas-clc/rest"
	"atlas-clc/services"
	"atlas-clc/session"
	"atlas-clc/socket/request"
	"atlas-clc/socket/request/handler"
	"atlas-clc/tasks"
	"context"
	"github.com/jtumidanski/atlas-socket"
	"github.com/sirupsen/logrus"
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

	lss := services.NewMapleSessionService(l)

	ss, err := socket.NewServer(l, lss, socket.IpAddress("0.0.0.0"), socket.Port(8484))
	if err != nil {
		return
	}

	registerHandlers(ss, l)
	go ss.Run()

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

	sessions := session.GetSessionRegistry().GetAll()
	for _, s := range sessions {
		lss.Destroy(s.SessionId())
	}

	l.Infoln("Service shutdown.")
}

func registerHandlers(ss *socket.Server, l logrus.FieldLogger) {
	hr := handlerRegister(ss, l)
	hr(handler.OpCodeLogin, &handler.LoginHandler{})
	hr(handler.OpCodeServerListReRequest, &handler.ServerListHandler{})
	hr(handler.OpCodeCharacterListWorld, &handler.CharacterListWorldHandler{})
	hr(handler.OpCodeServerStatus, &handler.ServerStatusHandler{})
	hr(handler.OpCodeServerRequest, &handler.ServerListHandler{})
	hr(handler.OpCodeClearWorldChannel, &handler.ClearWorldChannelHandler{})
	hr(handler.OpCodeCharacterListAll, &handler.CharacterListAllHandler{})
	hr(handler.OpCodeCharacterSelectFromAll, &handler.CharacterSelectFromAllHandler{})
	hr(handler.OpCodeCharacterSelectFromWorld, &handler.CharacterSelectFromWorldHandler{})
	hr(handler.OpCodeCharacterCheckName, &handler.CharacterCheckNameHandler{})
	hr(handler.OpCodeCharacterCreate, &handler.CharacterCreateHandler{})
	hr(handler.OpCodePong, &handler.PongHandler{})
	hr(handler.OpCodeClientStartError, &handler.ClientStartErrorHandler{})
}

func handlerRegister(ss *socket.Server, l logrus.FieldLogger) func(uint16, request.MapleHandler) {
	return func(op uint16, handler request.MapleHandler) {
		ss.RegisterHandler(op, request.AdaptHandler(l, handler))
	}
}
