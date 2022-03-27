package blocked_name

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func GetByName(l logrus.FieldLogger, span opentracing.Span) func(name string) (string, error) {
	return func(name string) (string, error) {
		dc, err := requestByName(name)(l, span)
		if err != nil {
			return "", err
		}

		if dc.Length() <= 0 {
			return "", err
		}
		var a = dc.Data().Attributes
		return a.Name, nil
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
