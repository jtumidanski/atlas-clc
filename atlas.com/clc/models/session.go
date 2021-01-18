package models

import (
   "atlas-clc/writers"
   "log"
   "math/rand"
   "net"
)

type Session struct {
   id   int
   con  net.Conn
   l    *log.Logger
   send AESOFB
   recv AESOFB
}

const (
   version uint16 = 83
)

func NewSession(id int, con *net.Conn, l *log.Logger) *Session {
   recvIv := []byte{70, 114, 122, 82}
   sendIv := []byte{82, 48, 120, 115}
   recvIv[3] = byte(rand.Float64() * 255)
   sendIv[3] = byte(rand.Float64() * 255)
   send := NewAESOFB(sendIv, uint16(65535)-version)
   recv := NewAESOFB(recvIv, version)
   return &Session{id, *con, l, *send, *recv}
}

func (s *Session) SessionId() int {
   return s.id
}

func (s *Session) Announce(m []byte) {
   _, err := s.con.Write(m)
   if err != nil {
      s.l.Fatal("[ERROR] Writing bytes to connection")
   }
}

func (r *Session) WriteHello() {
   r.Announce(writers.WriteHello(version, r.send.iv, r.recv.iv).Bytes())
}