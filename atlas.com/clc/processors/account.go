package processors

import (
	"atlas-clc/models"
	"atlas-clc/rest/requests"
	"log"
	"strconv"
)

func GetAccount(l *log.Logger, name string) (*models.Account, error) {
	a, err := requests.GetAccount(l, name)
	if err != nil {
		return nil, err
	}

	aid, err := strconv.Atoi(a.Data.Id)
	if err != nil {
		return nil, err
	}

	att := a.Data.Attributes
	return models.NewAccount(aid, att.Name, att.Password, att.Pin, att.Pic, att.LoggedIn, att.LastLogin, att.Gender, att.Banned, att.TOS, att.Language, att.Country, att.CharacterSlots), nil
}
