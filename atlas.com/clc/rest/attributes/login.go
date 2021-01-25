package attributes

type LoginListDataContainer struct {
   Data []LoginData `json:"data"`
}

type LoginDataContainer struct {
   Data LoginData `json:"data"`
}

type LoginInputDataContainer struct {
   Data LoginData `json:"data"`
}

type LoginData struct {
   Id         string          `json:"id"`
   Type       string          `json:"type"`
   Attributes LoginAttributes `json:"attributes"`
}

type LoginAttributes struct {
   SessionId int
   Name      string
   Password  string
   IpAddress string
   State     int
}
