package tasks

import (
	"atlas-clc/mapleSession"
	"atlas-clc/registries"
	"github.com/jtumidanski/atlas-socket/session"
	"github.com/sirupsen/logrus"
	"time"
)

type Timeout struct {
	l        logrus.FieldLogger
	lss      session.Service
	interval time.Duration
	timeout  time.Duration
}

func NewTimeout(l logrus.FieldLogger, lss session.Service, interval time.Duration) *Timeout {
	var to int64
	c, err := registries.GetConfiguration()
	if err != nil {
		to = 3600000
	} else {
		to = c.TimeoutDuration
	}

	timeout := time.Duration(to) * time.Millisecond
	l.Infof("Initializing timeout task to run every %dms, timeout session older than %dms", interval.Milliseconds(), timeout.Milliseconds())
	return &Timeout{l, lss, interval, timeout}
}

func (t *Timeout) Run() {
	sessions := registries.GetSessionRegistry().GetAll()
	cur := time.Now()

	for _, s := range sessions {
		as := s.(mapleSession.MapleSession)
		if cur.Sub(s.LastRequest()) > t.timeout {
			t.l.Infof("Account [%d] was auto-disconnected due to inactivity.", as.AccountId())
			t.lss.Destroy(as.SessionId())
		}
	}
}

func (t *Timeout) SleepTime() time.Duration {
	return t.interval
}
