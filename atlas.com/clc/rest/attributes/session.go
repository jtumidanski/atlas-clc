package attributes

type SessionListDataContainer struct {
   Data []SessionData `json:"data"`
}

type SessionDataContainer struct {
   Data SessionData `json:"data"`
}

type SessionInputDataContainer struct {
   Data SessionData `json:"data"`
}

type SessionData struct {
   Id         string            `json:"id"`
   Type       string            `json:"type"`
   Attributes SessionAttributes `json:"attributes"`
}

type SessionAttributes struct {
   AccountId int
}
