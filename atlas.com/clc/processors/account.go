package processors

import (
	"atlas-clc/models"
	"atlas-clc/rest/attributes"
	"atlas-clc/rest/requests"
	"log"
	"strconv"
)

func GetAccountByName(l *log.Logger, name string) (*models.Account, error) {
	a, err := requests.GetAccountByName(l, name)
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

func GetAccountById(l *log.Logger, id int) (*models.Account, error) {
	a, err := requests.GetAccountById(l, id)
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

func IsLoggedIn(l *log.Logger, id int) bool {
	a, err := GetAccountById(l,id)
	if err != nil {
		return false
	} else if a.LoggedIn() != 0 {
		return true
	} else {
		return false
	}
}

func makeAccount(aid int, att attributes.AccountAttributes) *models.Account {
	return models.NewAccountBuilder().
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
