package processors

import (
	"atlas-clc/rest/requests"
	"log"
)

func GetBlockedName(l *log.Logger, name string) (string, error) {
	a, err := requests.GetBlockedNamesByName(l, name)
	if err != nil {
		return "", err
	}
	if len(a.Data) <= 0 {
		return "", err
	}
	return a.Data[0].Attributes.Name, nil
}

func IsBlockedName(l *log.Logger, name string) (bool, error) {
	n, err := GetBlockedName(l, name)
	if err != nil {
		return true, err
	}
	if len(n) > 0 {
		return true, err
	}
	return false, err
}
