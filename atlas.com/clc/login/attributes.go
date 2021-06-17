package login

type inputDataContainer struct {
	Data dataBody `json:"data"`
}

type dataBody struct {
	Id         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes attributes `json:"attributes"`
}

type attributes struct {
	SessionId uint32 `json:"sessionId"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	IpAddress string `json:"ipAddress"`
	State     int    `json:"state"`
}
