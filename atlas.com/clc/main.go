package main

import (
	"atlas-clc/logger"
	"atlas-clc/registries"
	"atlas-clc/rest"
	"atlas-clc/services"
	"atlas-clc/socket/request"
	"atlas-clc/socket/request/handler"
	"atlas-clc/tasks"
	"github.com/jtumidanski/atlas-socket"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	l := logger.CreateLogger()

	config, err := registries.GetConfiguration()
	if err != nil {
		l.WithError(err).Fatal("Unable to successfully load configuration.")
	}

	lss := services.NewMapleSessionService(l)

	w := l.Writer()
	defer w.Close()

	ss, err := socket.NewServer(log.New(w, "", 0), lss, socket.IpAddress("0.0.0.0"), socket.Port(8484))
	if err != nil {
		return
	}

	registerHandlers(ss, l)
	go ss.Run()

	rs := rest.NewServer(l)
	go rs.Run()

	go tasks.Register(tasks.NewTimeout(l, lss, time.Second*time.Duration(config.TimeoutTaskInterval)))

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Infoln("Shutting down via signal:", sig)

	sessions := registries.GetSessionRegistry().GetAll()
	for _, s := range sessions {
		lss.Destroy(s.SessionId())
	}
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
