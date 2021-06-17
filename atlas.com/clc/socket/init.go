package socket

import (
	"atlas-clc/session"
	"atlas-clc/socket/request"
	"atlas-clc/socket/request/handler"
	"context"
	"github.com/jtumidanski/atlas-socket"
	request2 "github.com/jtumidanski/atlas-socket/request"
	"github.com/sirupsen/logrus"
	"sync"
)

func CreateSocketService(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		wg.Add(1)
		defer wg.Done()
		err := socket.Run(l, handlerProducer(l),
			socket.SetPort(8484),
			socket.SetSessionCreator(session.Create(l, session.GetRegistry())),
			socket.SetSessionMessageDecryptor(session.Decrypt(l, session.GetRegistry())),
			socket.SetSessionDestroyer(session.DestroyById(l, session.GetRegistry())),
		)
		if err != nil {
			l.WithError(err).Errorf("Socket service encountered error")
		}
	}()

	<-ctx.Done()
	l.Infof("Shutting down server on port 8484")
}

func handlerProducer(l logrus.FieldLogger) socket.MessageHandlerProducer {
	handlers := make(map[uint16]request2.Handler)
	handlers[handler.OpCodeLogin] = request.AdaptHandler(l, request.NoOpValidator, handler.HandleLoginRequest)
	handlers[handler.OpCodeServerListReRequest] = request.AdaptHandler(l, request.LoggedInValidator, handler.HandleServerListRequest)
	handlers[handler.OpCodeCharacterListWorld] = request.AdaptHandler(l, request.LoggedInValidator, handler.HandleCharacterListWorldRequest)
	handlers[handler.OpCodeServerStatus] = request.AdaptHandler(l, request.LoggedInValidator, handler.HandleServerStatusRequest)
	handlers[handler.OpCodeServerRequest] = request.AdaptHandler(l, request.LoggedInValidator, handler.HandleServerListRequest)
	handlers[handler.OpCodeClearWorldChannel] = request.AdaptHandler(l, request.LoggedInValidator, handler.HandleClearWorldChannelRequest)
	handlers[handler.OpCodeCharacterListAll] = request.AdaptHandler(l, request.LoggedInValidator, handler.HandleCharacterListAllRequest)
	handlers[handler.OpCodeCharacterSelectFromAll] = request.AdaptHandler(l, request.LoggedInValidator, handler.HandleCharacterSelectFromAllRequest)
	handlers[handler.OpCodeCharacterSelectFromWorld] = request.AdaptHandler(l, request.LoggedInValidator, handler.HandleCharacterSelectFromWorldRequest)
	handlers[handler.OpCodeCharacterCheckName] = request.AdaptHandler(l, request.LoggedInValidator, handler.HandleCheckCharacterNameRequest)
	handlers[handler.OpCodeCharacterCreate] = request.AdaptHandler(l, request.LoggedInValidator, handler.HandleCreateCharacterRequest)
	handlers[handler.OpCodePong] = request.AdaptHandler(l, request.LoggedInValidator, request.NoOpHandler)
	handlers[handler.OpCodeClientStartError] = request.AdaptHandler(l, request.NoOpValidator, handler.HandleClientStartErrorRequest)
	return func() map[uint16]request2.Handler {
		return handlers
	}
}
