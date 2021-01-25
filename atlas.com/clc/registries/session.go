package registries

import (
   "atlas-clc/sessions"
   "sync"
)

type SessionRegistry struct {
   mutex           sync.RWMutex
   sessionRegistry map[int]sessions.Session
}

var once sync.Once
var sessionRegistry *SessionRegistry

func GetSessionRegistry() *SessionRegistry {
   once.Do(func() {
      sessionRegistry = &SessionRegistry{}
      sessionRegistry.sessionRegistry = make(map[int]sessions.Session)
   })
   return sessionRegistry
}

func (r *SessionRegistry) AddSession(s *sessions.Session) {
   r.mutex.Lock()
   r.sessionRegistry[s.SessionId()] = *s
   r.mutex.Unlock()
}

func (r *SessionRegistry) GetSession(sessionId int) *sessions.Session {
   r.mutex.RLock()
   s := r.sessionRegistry[sessionId]
   r.mutex.RUnlock()
   return &s
}
