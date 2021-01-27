package requests

import (
	"atlas-clc/rest/attributes"
	"fmt"
	"log"
)

func GetBlockedNamesByName(l *log.Logger, name string) (*attributes.BlockedNameDataContainer, error) {
	ar := &attributes.BlockedNameDataContainer{}
	err := Get(l, fmt.Sprintf("http://atlas-nginx:80/ms/cos/blockedNames?name=%s", name), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
