package session

import "errors"

func Announce(b []byte) func(s Model) error {
	return func(s Model) error {
		if l, ok := GetRegistry().GetLock(s.SessionId()); ok {
			l.Lock()
			err := s.announceEncrypted(b)
			l.Unlock()
			return err
		}
		return errors.New("invalid session")
	}
}

func SetAccountId(accountId uint32) func(id uint32) Model {
	return func(id uint32) Model {
		s := Model{}
		var ok bool
		if s, ok = GetRegistry().Get(id); ok {
			s = s.setAccountId(accountId)
			GetRegistry().Update(s)
			return s
		}
		return s
	}
}

func UpdateLastRequest() func(id uint32) Model {
	return func(id uint32) Model {
		s := Model{}
		var ok bool
		if s, ok = GetRegistry().Get(id); ok {
			s = s.updateLastRequest()
			GetRegistry().Update(s)
			return s
		}
		return s
	}
}

func SetWorldId(worldId byte) func(id uint32) Model {
	return func(id uint32) Model {
		s := Model{}
		var ok bool
		if s, ok = GetRegistry().Get(id); ok {
			s = s.setWorldId(worldId)
			GetRegistry().Update(s)
			return s
		}
		return s
	}
}

func SetChannelId(channelId byte) func(id uint32) Model {
	return func(id uint32) Model {
		s := Model{}
		var ok bool
		if s, ok = GetRegistry().Get(id); ok {
			s = s.setChannelId(channelId)
			GetRegistry().Update(s)
			return s
		}
		return s
	}
}
