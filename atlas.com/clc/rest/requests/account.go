package requests

import (
	"atlas-clc/rest/attributes"
	"fmt"
	"log"
)

func GetAccount(l *log.Logger, name string) (*attributes.AccountDataContainer, error) {
	ar := &attributes.AccountDataContainer{}
	err := Get(l, fmt.Sprintf("http://atlas-nginx:80/ms/aos/accounts?name=%s", name), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
