package session

import (
	"atlas-clc/socket/response/writer"
	"github.com/jtumidanski/atlas-socket/crypto"
	"math/rand"
	"net"
	"time"
)

type Model struct {
	id         uint32
	accountId  uint32
	worldId    byte
	channelId  byte
	con        net.Conn
	send       crypto.AESOFB
	recv       crypto.AESOFB
	lastPacket time.Time
}

const (
	version uint16 = 83
)

func NewSession(id uint32, con net.Conn) *Model {
	recvIv := []byte{70, 114, 122, 82}
	sendIv := []byte{82, 48, 120, 115}
	recvIv[3] = byte(rand.Float64() * 255)
	sendIv[3] = byte(rand.Float64() * 255)
	send := crypto.NewAESOFB(sendIv, uint16(65535)-version)
	recv := crypto.NewAESOFB(recvIv, version)
	return &Model{id, 0, 0, 0, con, *send, *recv, time.Now()}
}

func (s *Model) SetAccountId(accountId uint32) {
	s.accountId = accountId
}

func (s *Model) SessionId() uint32 {
	return s.id
}

func (s *Model) AccountId() uint32 {
	return s.accountId
}

func (s *Model) Announce(b []byte) error {
	tmp := make([]byte, len(b)+4)
	copy(tmp, b)
	tmp = append([]byte{0, 0, 0, 0}, b...)
	tmp = s.send.Encrypt(tmp, true, true)
	_, err := s.con.Write(tmp)
	return err
}

func (s *Model) announce(b []byte) error {
	_, err := s.con.Write(b)
	return err
}

func (s *Model) WriteHello() {
	_ = s.announce(writer.WriteHello(nil)(version, s.send.IV(), s.recv.IV()))
}

func (s *Model) ReceiveAESOFB() *crypto.AESOFB {
	return &s.recv
}

func (s *Model) GetRemoteAddress() net.Addr {
	return s.con.RemoteAddr()
}

func (s *Model) SetWorldId(worldId byte) {
	s.worldId = worldId
}

func (s *Model) SetChannelId(channelId byte) {
	s.channelId = channelId
}

func (s *Model) WorldId() byte {
	return s.worldId
}

func (s *Model) ChannelId() byte {
	return s.channelId
}

func (s *Model) UpdateLastRequest() {
	s.lastPacket = time.Now()
}

func (s *Model) LastRequest() time.Time {
	return s.lastPacket
}

func (s *Model) Disconnect() {
	_ = s.con.Close()
}
