package socket

import (
	"atlas-clc/socket/crypto"
	"atlas-clc/socket/request"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type Server struct {
	logger               *log.Logger
	socketSessionService *SocketSessionService
	ipAddress            string
	port                 int
}

func NewServer(l *log.Logger, s *SocketSessionService, opts ...ServerOpt) (*Server, error) {
	server := Server{l, s, "0.0.0.0", 5000}
	for _, o := range opts {
		o(&server)
	}
	return &server, nil
}

func (s *Server) Run() {
	s.logger.Printf("[INFO] Starting tcp server on %s:%d", s.ipAddress, s.port)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.ipAddress, s.port))
	if err != nil {
		s.logger.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer lis.Close()

	sessionId := 0

	for {
		c, err := lis.Accept()
		if err != nil {
			s.logger.Println("Error connecting:", err.Error())
			return
		}
		s.logger.Println("[INFO] Client connected.")

		go newServerWorker(s.logger, c, sessionId, s.socketSessionService)
		sessionId += 1
	}
}

type serverWorker struct {
	logger               *log.Logger
	socketSessionService *SocketSessionService
}

type SocketSession interface {
	WriteHello()
}

type SocketSessionCreator interface {
	Create(int, *net.Conn, *log.Logger) (*SocketSession, error)
}

type SocketSessionStorer interface {
	Add(*SocketSession)
}

type SocketSessionDestroyer interface {
	Destroy(int)
}

type SocketSessionRetriever interface {
	Get(int) *SocketSession
	GetAll() []*SocketSession
}

type SocketSessionService interface {
	SocketSessionCreator
	SocketSessionStorer
	SocketSessionRetriever
	SocketSessionDestroyer
}

func newServerWorker(logger *log.Logger, conn net.Conn, sessionId int, service *SocketSessionService) *serverWorker {
	logger.Println("[INFO] Client " + conn.RemoteAddr().String() + " connected.")
	serverWorker := &serverWorker{logger, service}
	go serverWorker.Run(conn, sessionId, 4)
	return serverWorker
}

func (sw *serverWorker) Run(conn net.Conn, sessionId int, headerSize int) {
	s, err := sw.socketSessionService.Create(sessionId, conn, sw.logger)
	if err != nil {
		return
	}
	sw.socketSessionService.Add(s)

	s.WriteHello()

	header := true
	readSize := headerSize

	session := sw.socketSessionService.Get(sessionId)

	for {
		buffer := make([]byte, readSize)

		if _, err := conn.Read(buffer); err != nil {
			break
		}

		if header {
			readSize = crypto.GetPacketLength(buffer)
		} else {
			readSize = headerSize

			result := buffer
			if session.GetRecv() != nil {
				ue := session.GetRecv().Decrypt(buffer, true, true)
				result = ue
			}
			sw.handle(sessionId, result)
		}

		header = !header
	}

	Disconnect(sw.logger, session)

	sw.logger.Printf("[INFO] Session %d exiting read loop.", sessionId)
}

func (sw *serverWorker) handle(sessionId int, p request.Request) {
	go func(sessionId int, reader request.RequestReader) {
		op := reader.ReadUint16()
		h := request.GetHandle(sw.logger, op)
		if h != nil {
			h.Handle(sessionId, &reader)
		} else {
			sw.logger.Printf("[INFO] Session %d read a unhandled message with op %05X.", sessionId, op&0xFF)
		}
	}(sessionId, request.NewRequestReader(&p, time.Now().Unix()))
}
