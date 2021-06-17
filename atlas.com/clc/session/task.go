package session

import (
	"atlas-clc/configuration"
	"github.com/sirupsen/logrus"
	"time"
)

type Timeout struct {
	l        logrus.FieldLogger
	interval time.Duration
	timeout  time.Duration
}

func NewTimeout(l logrus.FieldLogger, interval time.Duration) *Timeout {
	var to int64
	c, err := configuration.GetConfiguration()
	if err != nil {
		to = 3600000
	} else {
		to = c.TimeoutDuration
	}

	timeout := time.Duration(to) * time.Millisecond
	l.Infof("Initializing timeout task to run every %dms, timeout session older than %dms", interval.Milliseconds(), timeout.Milliseconds())
	return &Timeout{l, interval, timeout}
}

func (t *Timeout) Run() {
	sessions := GetRegistry().GetAll()
	cur := time.Now()

	for _, s := range sessions {
		if cur.Sub(s.LastRequest()) > t.timeout {
			t.l.Infof("Account [%d] was auto-disconnected due to inactivity.", s.AccountId())
			DestroyById(t.l, GetRegistry())(s.SessionId())
		}
	}
}

func (t *Timeout) SleepTime() time.Duration {
	return t.interval
}
