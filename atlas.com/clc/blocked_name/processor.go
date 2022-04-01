package blocked_name

import (
	"atlas-clc/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func GetByName(l logrus.FieldLogger, span opentracing.Span) func(name string) (string, error) {
	return func(name string) (string, error) {
		return requests.Provider[attributes, string](l, span)(requestByName(name), getName)()
	}
}

func getName(body requests.DataBody[attributes]) (string, error) {
	attr := body.Attributes
	return attr.Name, nil
}

func IsBlockedName(l logrus.FieldLogger, span opentracing.Span) func(name string) (bool, error) {
	return func(name string) (bool, error) {
		n, err := GetByName(l, span)(name)
		if err != nil {
			return true, err
		}
		if len(n) > 0 {
			return true, err
		}
		return false, err
	}
}
