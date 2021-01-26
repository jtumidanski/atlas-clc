package models

type Account struct {
	id             int
	name           string
	password       string
	pin            string
	pic            string
	loggedIn       int
	lastLogin      uint64
	gender         byte
	banned         bool
	tos            bool
	language       string
	country        string
	characterSlots int16
}

func NewAccount(id int, name string, password string, pin string, pic string, loggedIn int, lastLogin uint64, gender byte, banned bool, tos bool, language string, country string, characterSlots int16) *Account {
	return &Account{
		id:             id,
		name:           name,
		password:       password,
		pin:            pin,
		pic:            pic,
		loggedIn:       loggedIn,
		lastLogin:      lastLogin,
		gender:         gender,
		banned:         banned,
		tos:            tos,
		language:       language,
		country:        country,
		characterSlots: characterSlots,
	}
}

func (a *Account) Id() int {
	return a.id
}

func (a *Account) Name() string {
	return a.name
}

func (a *Account) Gender() byte {
	return a.gender
}

func (a *Account) PIC() string {
	return a.pic
}

func (a *Account) CharacterSlots() int16 {
	return a.characterSlots
}
