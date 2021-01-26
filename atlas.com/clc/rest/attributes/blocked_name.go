package attributes

type BlockedNameListDataContainer struct {
	Data []BlockedNameData `json:"data"`
}

type BlockedNameDataContainer struct {
	Data BlockedNameData `json:"data"`
}

type BlockedNameInputDataContainer struct {
	Data BlockedNameData `json:"data"`
}

type BlockedNameData struct {
	Id         string                `json:"id"`
	Type       string                `json:"type"`
	Attributes BlockedNameAttributes `json:"attributes"`
}

type BlockedNameAttributes struct {
	Name string `json:"name"`
}
