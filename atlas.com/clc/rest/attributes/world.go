package attributes

type WorldListDataContainer struct {
	Data []WorldData `json:"data"`
}

type WorldDataContainer struct {
	Data WorldData `json:"data"`
}

type WorldInputDataContainer struct {
	Data WorldData `json:"data"`
}

type WorldData struct {
	Id         string          `json:"id"`
	Type       string          `json:"type"`
	Attributes WorldAttributes `json:"attributes"`
}

type WorldAttributes struct {
	Name               string `json:"name"`
	Flag               int    `json:"flag"`
	Message            string `json:"message"`
	EventMessage       string `json:"eventMessage"`
	Recommended        bool   `json:"recommended"`
	RecommendedMessage string `json:"recommendedMessage"`
	CapacityStatus     uint16    `json:"capacityStatus"`
}
