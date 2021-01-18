package registries

import (
   "atlas-clc/models"
   "sync"
)

type SessionRegistry struct {
   mutex           sync.RWMutex
   sessionRegistry map[int]models.Session
}

var once sync.Once
var sessionRegistry *SessionRegistry

func GetSessionRegistry() *SessionRegistry {
   once.Do(func() {
      sessionRegistry = &SessionRegistry{}
      sessionRegistry.sessionRegistry = make(map[int]models.Session)
   })
   return sessionRegistry
}

func (r *SessionRegistry) AddSession(s *models.Session) {
   r.mutex.Lock()
   r.sessionRegistry[s.SessionId()] = *s
   r.mutex.Unlock()
}