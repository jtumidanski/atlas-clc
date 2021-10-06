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
				socket.SetSessionDestroyer(session.DestroyByIdWithSpan(l, session.GetRegistry())),
			)
			if err != nil {
				l.WithError(err).Errorf("Socket service encountered error")
			}
		}()

		<-ctx.Done()
		l.Infof("Shutting down server on port 8484")
	}()
}

const (
	LoginRequest = "login_request"
	ServerListRequest = "server_list_request"
	CharacterListWorldRequest = "character_list_world_request"
	ServerStatus = "server_status"
	ServerRequest = "server_request"
	ClearWorldChannel = "clear_world_channel"
	CharacterListAll = "character_list_all"
	CharacterSelectFromAll = "character_select_from_all"
	CharacterSelectFromWorld = "character_select_from_world"
	CharacterCheckName = "character_check_name"
	CharacterCreate = "character_create"
	Pong = "pong"
	ClientStartError = "client_start_error"
)

func handlerProducer(l logrus.FieldLogger) socket.MessageHandlerProducer {
	handlers := make(map[uint16]request2.Handler)
	hr := func(op uint16, name string, v request.MessageValidator, h request.MessageHandler) {
		handlers[op] = request.AdaptHandler(l, name, v, h)
	}

	hr(handler.OpCodeLogin, LoginRequest, request.NoOpValidator, handler.HandleLoginRequest)
	hr(handler.OpCodeServerListReRequest, ServerListRequest, request.LoggedInValidator, handler.HandleServerListRequest)
	hr(handler.OpCodeCharacterListWorld, CharacterListWorldRequest, request.LoggedInValidator, handler.HandleCharacterListWorldRequest)
	hr(handler.OpCodeServerStatus, ServerStatus, request.LoggedInValidator, handler.HandleServerStatusRequest)
	hr(handler.OpCodeServerRequest, ServerRequest, request.LoggedInValidator, handler.HandleServerListRequest)
	hr(handler.OpCodeClearWorldChannel, ClearWorldChannel, request.LoggedInValidator, handler.HandleClearWorldChannelRequest)
	hr(handler.OpCodeCharacterListAll, CharacterListAll, request.LoggedInValidator, handler.HandleCharacterListAllRequest)
	hr(handler.OpCodeCharacterSelectFromAll, CharacterSelectFromAll, request.LoggedInValidator, handler.HandleCharacterSelectFromAllRequest)
	hr(handler.OpCodeCharacterSelectFromWorld, CharacterSelectFromWorld, request.LoggedInValidator, handler.HandleCharacterSelectFromWorldRequest)
	hr(handler.OpCodeCharacterCheckName, CharacterCheckName, request.LoggedInValidator, handler.HandleCheckCharacterNameRequest)
	hr(handler.OpCodeCharacterCreate, CharacterCreate, request.LoggedInValidator, handler.HandleCreateCharacterRequest)
	hr(handler.OpCodePong, Pong, request.LoggedInValidator, request.NoOpHandler)
	hr(handler.OpCodeClientStartError, ClientStartError, request.NoOpValidator, handler.HandleClientStartErrorRequest)

	return func() map[uint16]request2.Handler {
		return handlers
	}
}
