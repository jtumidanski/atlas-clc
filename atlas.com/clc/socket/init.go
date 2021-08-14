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
	go func() {
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
	}()
}

func handlerProducer(l logrus.FieldLogger) socket.MessageHandlerProducer {
	handlers := make(map[uint16]request2.Handler)
	hr := func(op uint16, v request.MessageValidator, h request.MessageHandler) {
		handlers[op] = request.AdaptHandler(l, v, h)
	}

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

	return func() map[uint16]request2.Handler {
		return handlers
	}
}
