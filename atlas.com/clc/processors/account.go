package processors

import (
	"atlas-clc/domain"
	"atlas-clc/rest/attributes"
	"atlas-clc/rest/requests"
	"strconv"
)

func GetAccountByName(name string) (*domain.Account, error) {
	a, err := requests.GetAccountByName(name)
	if err != nil {
		return nil, err
	}

	d := a.Data()
	aid, err := strconv.Atoi(d.Id)
	if err != nil {
		return nil, err
	}

	return makeAccount(aid, d.Attributes), nil
}

func GetAccountById(id int) (*domain.Account, error) {
	a, err := requests.GetAccountById(id)
	if err != nil {
		return nil, err
	}

	d := a.Data()
	aid, err := strconv.Atoi(d.Id)
	if err != nil {
		return nil, err
	}

	return makeAccount(aid, d.Attributes), nil
}

func IsLoggedIn(id int) bool {
	a, err := GetAccountById(id)
	if err != nil {
		return false
	} else if a.LoggedIn() != 0 {
		return true
	} else {
		return false
	}
}

func makeAccount(aid int, att attributes.AccountAttributes) *domain.Account {
	return domain.NewAccountBuilder().
		SetId(aid).
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
}
