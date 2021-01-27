package processors

import (
	"atlas-clc/models"
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

	att := d.Attributes
	return models.NewAccount(aid, att.Name, att.Password, att.Pin, att.Pic, att.LoggedIn, att.LastLogin, att.Gender, att.Banned, att.TOS, att.Language, att.Country, att.CharacterSlots), nil
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

	att := d.Attributes
	return models.NewAccount(aid, att.Name, att.Password, att.Pin, att.Pic, att.LoggedIn, att.LastLogin, att.Gender, att.Banned, att.TOS, att.Language, att.Country, att.CharacterSlots), nil
}
