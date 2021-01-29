package sessions

import (
	"atlas-clc/socket/response/writer"
	"github.com/jtumidanski/atlas-socket/crypto"
	session2 "github.com/jtumidanski/atlas-socket/session"
	"log"
	"math/rand"
	"net"
	"time"
)

type Session interface {
	session2.Session
	SessionId() int
	Disconnect()
	AccountId() int
}

type session struct {
	id         int
	accountId  int
	worldId    byte
	channelId  byte
	con        net.Conn
	l          *log.Logger
	send       crypto.AESOFB
	recv       crypto.AESOFB
	lastPacket time.Time
}

const (
	version uint16 = 83
)

func NewSession(id int, con net.Conn, l *log.Logger) Session {
	recvIv := []byte{70, 114, 122, 82}
	sendIv := []byte{82, 48, 120, 115}
	recvIv[3] = byte(rand.Float64() * 255)
	sendIv[3] = byte(rand.Float64() * 255)
	send := crypto.NewAESOFB(sendIv, uint16(65535)-version)
	recv := crypto.NewAESOFB(recvIv, version)
	return &session{id, -1, 0, 0, con, l, *send, *recv, time.Now()}
}

func (s *session) SetAccountId(accountId int) {
	s.accountId = accountId
}

func (s *session) SessionId() int {
	return s.id
}

func (s *session) AccountId() int {
	return s.accountId
}

func (s *session) Announce(b []byte) {
	tmp := make([]byte, len(b)+4)
	copy(tmp, b)
	tmp = append([]byte{0, 0, 0, 0}, b...)
	tmp = s.send.Encrypt(tmp, true, true)
	_, err := s.con.Write(tmp)
	if err != nil {
		s.l.Fatal("[ERROR] Writing bytes to connection")
	}
}

func (s *session) announce(b []byte) {
	_, err := s.con.Write(b)
	if err != nil {
		s.l.Fatal("[ERROR] Writing bytes to connection")
	}
}

func (s *session) WriteHello() {
	s.announce(writer.WriteHello(version, s.send.IV(), s.recv.IV()))
}

func (s *session) ReceiveAESOFB() *crypto.AESOFB {
	return &s.recv
}

func (s *session) GetRemoteAddress() net.Addr {
	return s.con.RemoteAddr()
}

func (s *session) SetWorldId(worldId byte) {
	s.worldId = worldId
}

func (s *session) SetChannelId(channelId byte) {
	s.channelId = channelId
}

func (s *session) WorldId() byte {
	return s.worldId
}

func (s *session) ChannelId() byte {
	return s.channelId
}

func (s *session) UpdateLastPacket() {
	s.lastPacket = time.Now()
}

func (s *session) LastPacket() time.Time {
	return s.lastPacket
}

func (s *session) Disconnect() {
	_ = s.con.Close()
}
