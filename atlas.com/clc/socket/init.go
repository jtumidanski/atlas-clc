package socket

import (
	"atlas-clc/session"
	"atlas-clc/socket/request"
	"atlas-clc/socket/request/handler"
	"context"
	"github.com/jtumidanski/atlas-socket"
	"github.com/sirupsen/logrus"
	"sync"
)

func CreateSocketService(l *logrus.Logger, s session.Service, ctx context.Context, wg *sync.WaitGroup) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		wg.Add(1)
		defer wg.Done()

		ss, err := socket.NewServer(l, s, socket.IpAddress("0.0.0.0"), socket.Port(8484))
		if err != nil {
			return
		}
		registerHandlers(ss, l)
		ss.Run()
	}()

	<-ctx.Done()
	l.Infof("Shutting down server on port 8484")
}

func registerHandlers(ss *socket.Server, l logrus.FieldLogger) {
	hr := handlerRegister(ss, l)
	hr(handler.OpCodeLogin, request.NoOpValidator, handler.HandleLoginRequest)
	hr(handler.OpCodeServerListReRequest, request.LoggedInValidator, handler.HandleServerListRequest)
	hr(handler.OpCodeCharacterListWorld, request.LoggedInValidator, handler.HandleCharacterListWorldRequest)
	hr(handler.OpCodeServerStatus, request.LoggedInValidator, handler.HandleServerStatusRequest)
	hr(handler.OpCodeServerRequest, request.LoggedInValidator, handler.HandleServerListRequest)
	hr(handler.OpCodeClearWorldChannel, request.LoggedInValidator, handler.HandleClearWorldChannelRequest)
	hr(handler.OpCodeCharacterListAll, request.LoggedInValidator, handler.HandleCharacterListAllRequest)
	hr(handler.OpCodeCharacterSelectFromAll, request.LoggedInValidator, handler.HandleCharacterSelectFromAllRequest)
	hr(handler.OpCodeCharacterSelectFromWorld, request.LoggedInValidator, handler.HandleCharacterSelectFromWorldRequest)
	hr(handler.OpCodeCharacterCheckName, request.LoggedInValidator, handler.HandleCheckCharacterNameRequest)
	hr(handler.OpCodeCharacterCreate, request.LoggedInValidator, handler.HandleCreateCharacterRequest)
	hr(handler.OpCodePong, request.LoggedInValidator, request.NoOpHandler)
	hr(handler.OpCodeClientStartError, request.NoOpValidator, handler.HandleClientStartErrorRequest)
}

func handlerRegister(ss *socket.Server, l logrus.FieldLogger) func(uint16, request.MessageValidator, request.MessageHandler) {
	return func(op uint16, v request.MessageValidator, h request.MessageHandler) {
		ss.RegisterHandler(op, request.AdaptHandler(l, v, h))
	}
}
