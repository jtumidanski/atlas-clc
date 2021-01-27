package attributes

type LoginInputDataContainer struct {
	Data LoginData `json:"data"`
}

type LoginData struct {
	Id         string          `json:"id"`
	Type       string          `json:"type"`
	Attributes LoginAttributes `json:"attributes"`
}

type LoginAttributes struct {
	SessionId int    `json:"sessionId"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	IpAddress string `json:"ipAddress"`
	State     int    `json:"state"`
}
