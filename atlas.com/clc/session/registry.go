package session

import (
	"sync"
)

type Registry struct {
	mutex           sync.RWMutex
	sessionRegistry map[int]*MapleSession
}

var sessionRegistryOnce sync.Once
var sessionRegistry *Registry

func GetRegistry() *Registry {
	sessionRegistryOnce.Do(func() {
		sessionRegistry = &Registry{}
		sessionRegistry.sessionRegistry = make(map[int]*MapleSession)
	})
	return sessionRegistry
}

func (r *Registry) Add(s *MapleSession) {
	r.mutex.Lock()
	r.sessionRegistry[(*s).SessionId()] = s
	r.mutex.Unlock()
}

func (r *Registry) Remove(sessionId int) {
	r.mutex.Lock()
	delete(r.sessionRegistry, sessionId)
	r.mutex.Unlock()
}

func (r *Registry) Get(sessionId int) MapleSession {
	r.mutex.RLock()
	s := r.sessionRegistry[sessionId]
	r.mutex.RUnlock()
	return *s
}

func (r *Registry) GetAll() []MapleSession {
	r.mutex.RLock()
	s := make([]MapleSession, 0)
	for _, v := range r.sessionRegistry {
		s = append(s, *v)
	}
	r.mutex.RUnlock()
	return s
}
