package session

import (
	"sync"
)

type SessionRegistry struct {
	mutex           sync.RWMutex
	sessionRegistry map[int]*MapleSession
}

var sessionRegistryOnce sync.Once
var sessionRegistry *SessionRegistry

func GetSessionRegistry() *SessionRegistry {
	sessionRegistryOnce.Do(func() {
		sessionRegistry = &SessionRegistry{}
		sessionRegistry.sessionRegistry = make(map[int]*MapleSession)
	})
	return sessionRegistry
}

func (r *SessionRegistry) Add(s *MapleSession) {
	r.mutex.Lock()
	r.sessionRegistry[(*s).SessionId()] = s
	r.mutex.Unlock()
}

func (r *SessionRegistry) Remove(sessionId int) {
	r.mutex.Lock()
	delete(r.sessionRegistry, sessionId)
	r.mutex.Unlock()
}

func (r *SessionRegistry) Get(sessionId int) MapleSession {
	r.mutex.RLock()
	s := r.sessionRegistry[sessionId]
	r.mutex.RUnlock()
	return *s
}

func (r *SessionRegistry) GetAll() []MapleSession {
	r.mutex.RLock()
	s := make([]MapleSession, 0)
	for _, v := range r.sessionRegistry {
		s = append(s, *v)
	}
	r.mutex.RUnlock()
	return s
}
