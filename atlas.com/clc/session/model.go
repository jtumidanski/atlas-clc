package session

import (
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

func NewSession(id uint32, con net.Conn) Model {
	recvIv := []byte{70, 114, 122, 82}
	sendIv := []byte{82, 48, 120, 115}
	recvIv[3] = byte(rand.Float64() * 255)
	sendIv[3] = byte(rand.Float64() * 255)
	send := crypto.NewAESOFB(sendIv, uint16(65535)-version)
	recv := crypto.NewAESOFB(recvIv, version)
	return Model{id, 0, 0, 0, con, *send, *recv, time.Now()}
}

func CloneSession(s Model) Model {
	return Model{
		id:         s.id,
		accountId:  s.accountId,
		worldId:    s.worldId,
		channelId:  s.channelId,
		con:        s.con,
		send:       s.send,
		recv:       s.recv,
		lastPacket: s.lastPacket,
	}
}

func (s *Model) setAccountId(accountId uint32) Model {
	ns := CloneSession(*s)
	ns.accountId = accountId
	return ns
}

func (s *Model) SessionId() uint32 {
	return s.id
}

func (s *Model) AccountId() uint32 {
	return s.accountId
}

func (s *Model) announceEncrypted(b []byte) error {
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

func (s *Model) WriteHello() error {
	return s.announce(WriteHello(nil)(version, s.send.IV(), s.recv.IV()))
}

func (s *Model) ReceiveAESOFB() *crypto.AESOFB {
	return &s.recv
}

func (s *Model) GetRemoteAddress() net.Addr {
	return s.con.RemoteAddr()
}

func (s *Model) setWorldId(worldId byte) Model {
	ns := CloneSession(*s)
	ns.worldId = worldId
	return ns
}

func (s *Model) setChannelId(channelId byte) Model {
	ns := CloneSession(*s)
	ns.channelId = channelId
	return ns
}

func (s *Model) WorldId() byte {
	return s.worldId
}

func (s *Model) ChannelId() byte {
	return s.channelId
}

func (s *Model) updateLastRequest() Model {
	ns := CloneSession(*s)
	ns.lastPacket = time.Now()
	return ns
}

func (s *Model) LastRequest() time.Time {
	return s.lastPacket
}

func (s *Model) Disconnect() {
	_ = s.con.Close()
}
