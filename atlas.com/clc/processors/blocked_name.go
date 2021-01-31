package processors

import (
	"atlas-clc/rest/requests"
)

func GetBlockedName(name string) (string, error) {
	a, err := requests.GetBlockedNamesByName(name)
	if err != nil {
		return "", err
	}
	if len(a.DataList()) <= 0 {
		return "", err
	}
	return a.Data().Attributes.Name, nil
}

func IsBlockedName(name string) (bool, error) {
	n, err := GetBlockedName(name)
	if err != nil {
		return true, err
	}
	if len(n) > 0 {
		return true, err
	}
	return false, err
}
