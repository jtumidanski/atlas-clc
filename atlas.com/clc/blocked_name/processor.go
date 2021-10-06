package blocked_name

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func GetByName(l logrus.FieldLogger, span opentracing.Span) func(name string) (string, error) {
	return func(name string) (string, error) {
		a, err := requestByName(l, span)(name)
		if err != nil {
			return "", err
		}
		if len(a.DataList()) <= 0 {
			return "", err
		}
		return a.Data().Attributes.Name, nil
	}
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
