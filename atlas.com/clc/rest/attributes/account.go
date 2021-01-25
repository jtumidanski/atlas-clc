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
	Name           string
	Password       string
	Pin            string
	Pic            string
	LoggedIn       int
	LastLogin      string
	Gender         byte
	Banned         bool
	TOS            bool
	Language       string
	Country        string
	CharacterSlots int16
}
