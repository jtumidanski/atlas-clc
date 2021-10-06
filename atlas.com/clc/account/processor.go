package account

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"strconv"
)

type ModelOperator func(*Model)

type ModelProvider func() (*Model, error)

type ModelListProvider func() ([]*Model, error)

func requestModelProvider(l logrus.FieldLogger, span opentracing.Span) func(r Request) ModelProvider {
	return func(r Request) ModelProvider {
		return func() (*Model, error) {
			resp, err := r(l, span)
			if err != nil {
				return nil, err
			}

			p, err := makeModel(resp.Data())
			if err != nil {
				return nil, err
			}
			return p, nil
		}
	}
}

func For(provider ModelProvider, operator ModelOperator) {
	m, err := provider()
	if err != nil {
		return
	}
	operator(m)
}

func ForAccountByName(l logrus.FieldLogger, span opentracing.Span) func(name string, operator ModelOperator) {
	return func(name string, operator ModelOperator) {
		For(ByNameModelProvider(l, span)(name), operator)
	}
}

func ByNameModelProvider(l logrus.FieldLogger, span opentracing.Span) func(name string) ModelProvider {
	return func(name string) ModelProvider {
		return requestModelProvider(l, span)(requestAccountByName(name))
	}
}

func ByIdModelProvider(l logrus.FieldLogger, span opentracing.Span) func(id uint32) ModelProvider {
	return func(id uint32) ModelProvider {
		return requestModelProvider(l, span)(requestAccountById(id))
	}
}

func GetById(l logrus.FieldLogger, span opentracing.Span) func(id uint32) (*Model, error) {
	return func(id uint32) (*Model, error) {
		return ByIdModelProvider(l, span)(id)()
	}
}

func IsLoggedIn(l logrus.FieldLogger, span opentracing.Span) func(id uint32) bool {
	return func(id uint32) bool {
		a, err := GetById(l, span)(id)
		if err != nil {
			return false
		} else if a.LoggedIn() != 0 {
			return true
		} else {
			return false
		}
	}
}

func makeModel(body *dataBody) (*Model, error) {
	id, err := strconv.ParseUint(body.Id, 10, 32)
	if err != nil {
		return nil, err
	}
	att := body.Attributes
	m := NewBuilder().
		SetId(uint32(id)).
		SetPassword(att.Password).
		SetPin(att.Pin).
		SetPic(att.Pic).
		SetLoggedIn(att.LoggedIn).
		SetLastLogin(att.LastLogin).
		SetGender(att.Gender).
		SetBanned(att.Banned).
		SetTos(att.TOS).
		SetLanguage(att.Language).
		SetCountry(att.Country).
		SetCharacterSlots(att.CharacterSlots).
		Build()
	return &m, nil
}
