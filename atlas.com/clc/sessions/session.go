package sessions

import (
	"atlas-clc/crypto"
	"atlas-clc/packets/outputs/writers"
	"log"
	"math/rand"
	"net"
)

type Session struct {
	id        int
	accountId int
	con       net.Conn
	l         *log.Logger
	send      crypto.AESOFB
	recv      crypto.AESOFB
}

const (
	version uint16 = 83
)

func NewSession(id int, con *net.Conn, l *log.Logger) *Session {
	recvIv := []byte{70, 114, 122, 82}
	sendIv := []byte{82, 48, 120, 115}
	recvIv[3] = byte(rand.Float64() * 255)
	sendIv[3] = byte(rand.Float64() * 255)
	send := crypto.NewAESOFB(sendIv, uint16(65535)-version)
	recv := crypto.NewAESOFB(recvIv, version)
	return &Session{id, -1, *con, l, *send, *recv}
}

func (s *Session) SetAccountId(accountId int) {
	s.accountId = accountId
}

func (s *Session) SessionId() int {
	return s.id
}

func (s *Session) Announce(b []byte) {
	tmp := make([]byte, len(b)+4)
	copy(tmp, b)
	tmp = append([]byte{0, 0, 0, 0}, b...)
	tmp = s.send.Encrypt(tmp, true, true)
	_, err := s.con.Write(tmp)
	if err != nil {
		s.l.Fatal("[ERROR] Writing bytes to connection")
	}
}

func (s *Session) announce(b []byte) {
	_, err := s.con.Write(b)
	if err != nil {
		s.l.Fatal("[ERROR] Writing bytes to connection")
	}
}

func (s *Session) WriteHello() {
	s.announce(writers.WriteHello(version, s.send.IV(), s.recv.IV()))
}

func (s *Session) GetRecv() *crypto.AESOFB {
	return &s.recv
}

func (s *Session) GetRemoteAddress() net.Addr {
	return s.con.RemoteAddr()
}
