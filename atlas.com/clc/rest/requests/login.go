package requests

import (
   "atlas-clc/rest/attributes"
   "bytes"
   "encoding/json"
   "fmt"
   "log"
   "net/http"
)

func CreateLogin(l *log.Logger, sessionId int, name string, password string, ipAddress string) (int, *attributes.LoginData, *attributes.ErrorListDataContainer) {
   i := attributes.LoginInputDataContainer{
      Data: attributes.LoginData{
         Id:   "0",
         Type: "com.atlas.aos.attribute.LoginAttributes",
         Attributes: attributes.LoginAttributes{
            SessionId: sessionId,
            Name:      name,
            Password:  password,
            IpAddress: ipAddress,
            State:     0,
         },
      },
   }

   jsonReq, err := json.Marshal(i)
   if err != nil {
      l.Fatal("[ERROR] marshalling [LoginAttributes]")
      return 500, nil, nil
   }

   r, err := http.Post(fmt.Sprintf("http://atlas-nginx:80/ms/aos/logins"),
      "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
   if err != nil {
      l.Printf("[ERROR] dispatching login for %s", name)
      return 500, nil, nil
   }

   if r.StatusCode != http.StatusNoContent {
      rb := &attributes.ErrorListDataContainer{}
      err = attributes.FromJSON(rb, r.Body)
      if err != nil {
         l.Printf("[ERROR] decoding error response")
         return 500, nil, nil
      }

      return r.StatusCode, nil, rb
   } else {
      rb := &attributes.LoginDataContainer{}
      err = attributes.FromJSON(rb, r.Body)
      if err != nil {
         l.Printf("[ERROR] decoding login response")
         return 500, nil, nil
      }

      return r.StatusCode, &rb.Data, nil
   }
}
