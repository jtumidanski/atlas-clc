package models

type LoginPassword struct {
   login    string
   password string
   hwid     []byte
}

func NewLoginPassword(login string, password string, hwid []byte) *LoginPassword {
   return &LoginPassword{
      login:    login,
      password: password,
      hwid:     hwid,
   }
}

func (l *LoginPassword) Login() string {
   return l.login
}

func (l *LoginPassword) Password() string {
   return l.password
}
