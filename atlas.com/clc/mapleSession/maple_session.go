package mapleSession

import (
	"atlas-clc/socket/response/writer"
	"github.com/jtumidanski/atlas-socket/crypto"
	"github.com/jtumidanski/atlas-socket/session"
	"log"
	"math/rand"
	"net"
	"time"
)

type AccountSession interface {
	AccountId() int
	SetAccountId(id int)
}

type MapleSession interface {
	session.Session
	AccountSession
	Announce(response []byte)
	WorldId() byte
	SetWorldId(id byte)
	SetChannelId(id byte)
	ChannelId() byte
}

type mapleSession struct {
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

func NewSession(id int, con net.Conn, l *log.Logger) MapleSession {
	recvIv := []byte{70, 114, 122, 82}
	sendIv := []byte{82, 48, 120, 115}
	recvIv[3] = byte(rand.Float64() * 255)
	sendIv[3] = byte(rand.Float64() * 255)
	send := crypto.NewAESOFB(sendIv, uint16(65535)-version)
	recv := crypto.NewAESOFB(recvIv, version)
	return &mapleSession{id, -1, 0, 0, con, l, *send, *recv, time.Now()}
}

func (s *mapleSession) SetAccountId(accountId int) {
	s.accountId = accountId
}

func (s *mapleSession) SessionId() int {
	return s.id
}

func (s *mapleSession) AccountId() int {
	return s.accountId
}

func (s *mapleSession) Announce(b []byte) {
	tmp := make([]byte, len(b)+4)
	copy(tmp, b)
	tmp = append([]byte{0, 0, 0, 0}, b...)
	tmp = s.send.Encrypt(tmp, true, true)
	_, err := s.con.Write(tmp)
	if err != nil {
		s.l.Fatal("[ERROR] Writing bytes to connection")
	}
}

func (s *mapleSession) announce(b []byte) {
	_, err := s.con.Write(b)
	if err != nil {
		s.l.Fatal("[ERROR] Writing bytes to connection")
	}
}

func (s *mapleSession) WriteHello() {
	s.announce(writer.WriteHello(version, s.send.IV(), s.recv.IV()))
}

func (s *mapleSession) ReceiveAESOFB() *crypto.AESOFB {
	return &s.recv
}

func (s *mapleSession) GetRemoteAddress() net.Addr {
	return s.con.RemoteAddr()
}

func (s *mapleSession) SetWorldId(worldId byte) {
	s.worldId = worldId
}

func (s *mapleSession) SetChannelId(channelId byte) {
	s.channelId = channelId
}

func (s *mapleSession) WorldId() byte {
	return s.worldId
}

func (s *mapleSession) ChannelId() byte {
	return s.channelId
}

func (s *mapleSession) UpdateLastRequest() {
	s.lastPacket = time.Now()
}

func (s *mapleSession) LastRequest() time.Time {
	return s.lastPacket
}

func (s *mapleSession) Disconnect() {
	_ = s.con.Close()
}
