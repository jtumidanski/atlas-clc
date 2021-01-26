package requests

import (
	"atlas-clc/rest/attributes"
	"fmt"
	"log"
)

func GetBlockedNamesByName(l *log.Logger, name string) (*attributes.BlockedNameListDataContainer, error) {
	ar := &attributes.BlockedNameListDataContainer{}
	err := Get(l, fmt.Sprintf("http://atlas-nginx:80/ms/cos/blockedNames?name=%s", name), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
