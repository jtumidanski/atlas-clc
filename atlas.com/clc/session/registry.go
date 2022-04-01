package session

import (
	"sync"
)

type Registry struct {
	mutex           sync.RWMutex
	sessionRegistry map[uint32]Model
	lockRegistry    map[uint32]*sync.RWMutex
}

var sessionRegistryOnce sync.Once
var sessionRegistry *Registry

func GetRegistry() *Registry {
	sessionRegistryOnce.Do(func() {
		sessionRegistry = &Registry{}
		sessionRegistry.sessionRegistry = make(map[uint32]Model)
		sessionRegistry.lockRegistry = make(map[uint32]*sync.RWMutex)
	})
	return sessionRegistry
}

func (r *Registry) Add(s Model) {
	r.mutex.Lock()
	r.sessionRegistry[s.SessionId()] = s
	r.lockRegistry[s.SessionId()] = &sync.RWMutex{}
	r.mutex.Unlock()
}

func (r *Registry) Remove(sessionId uint32) {
	r.mutex.Lock()
	delete(r.sessionRegistry, sessionId)
	delete(r.lockRegistry, sessionId)
	r.mutex.Unlock()
}

func (r *Registry) Get(sessionId uint32) (Model, bool) {
	r.mutex.RLock()
	if s, ok := r.sessionRegistry[sessionId]; ok {
		r.mutex.RUnlock()
		return s, true
	}
	r.mutex.RUnlock()
	return Model{}, false
}

func (r *Registry) GetLock(sessionId uint32) (*sync.RWMutex, bool) {
	r.mutex.RLock()
	if val, ok := r.lockRegistry[sessionId]; ok {
		r.mutex.RUnlock()
		return val, true
	}
	r.mutex.RUnlock()
	return nil, false
}

func (r *Registry) GetAll() []Model {
	r.mutex.RLock()
	s := make([]Model, 0)
	for _, v := range r.sessionRegistry {
		s = append(s, v)
	}
	r.mutex.RUnlock()
	return s
}

func (r *Registry) Update(m Model) {
	r.mutex.Lock()
	if _, ok := r.sessionRegistry[m.SessionId()]; ok {
		r.sessionRegistry[m.SessionId()] = m
	}
	r.mutex.Unlock()
}
