package attributes

type AccountListDataContainer struct {
	Data []AccountData `json:"data"`
}

type AccountDataContainer struct {
	Data AccountData `json:"data"`
}

type AccountInputDataContainer struct {
	Data AccountData `json:"data"`
}

type AccountData struct {
	Id         string            `json:"id"`
	Type       string            `json:"type"`
	Attributes AccountAttributes `json:"attributes"`
}

type AccountAttributes struct {
	Name           string `json:"name"`
	Password       string `json:"password"`
	Pin            string `json:"pin"`
	Pic            string `json:"pic"`
	LoggedIn       int    `json:"loggedIn"`
	LastLogin      uint64 `json:"lastLogin"`
	Gender         byte   `json:"gender"`
	Banned         bool   `json:"banned"`
	TOS            bool   `json:"tos"`
	Language       string `json:"language"`
	Country        string `json:"country"`
	CharacterSlots int16  `json:"characterSlots"`
}
