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

func (r *SessionRegistry) RemoveSession(sessionId int) {
   r.mutex.Lock()
   delete(r.sessionRegistry, sessionId)
   r.mutex.Unlock()
}

func (r *SessionRegistry) GetSession(sessionId int) *sessions.Session {
   var ses *sessions.Session
   r.mutex.RLock()
   if val, ok := r.sessionRegistry[sessionId]; ok {
      ses = &val
   }
   r.mutex.RUnlock()
   return ses
}

func (r *SessionRegistry) GetSessions() []sessions.Session {
   r.mutex.RLock()
   s := make([]sessions.Session, 0)
   for _, v := range r.sessionRegistry {
      s = append(s, v)
   }
   r.mutex.RUnlock()
   return s
}
