package character

import (
	"atlas-clc/blocked_name"
	"atlas-clc/inventory"
	"atlas-clc/pet"
	"errors"
	"regexp"
	"strconv"
)

func GetPropertiesByName(name string) (*Properties, error) {
	ca, err := requestPropertiesByName(name)
	if err != nil {
		return nil, err
	}
	if len(ca.DataList()) <= 0 {
		return nil, errors.New("unable to find character by name")
	}

	return makeProperties(ca.Data()), nil
}

func makeProperties(ca *propertiesDataBody) *Properties {
	cid, err := strconv.ParseUint(ca.Id, 10, 32)
	if err != nil {
		return nil
	}
	att := ca.Attributes
	r := NewCharacterPropertiesBuilder().
		SetId(uint32(cid)).
		SetWorldId(att.WorldId).
		SetName(att.Name).
		SetGender(att.Gender).
		SetSkinColor(att.SkinColor).
		SetFace(att.Face).
		SetHair(att.Hair).
		SetLevel(att.Level).
		SetJobId(att.JobId).
		SetStrength(att.Strength).
		SetDexterity(att.Dexterity).
		SetIntelligence(att.Intelligence).
		SetLuck(att.Luck).
		SetHp(att.Hp).
		SetMaxHp(att.MaxHp).
		SetMp(att.Mp).
		SetMaxMp(att.MaxMp).
		SetAp(att.Ap).
		SetSp(att.Sp).
		SetExperience(att.Experience).
		SetFame(att.Fame).
		SetGachaponExperience(att.GachaponExperience).
		SetMapId(att.MapId).
		SetSpawnPoint(att.SpawnPoint).
		Build()
	return &r
}

func IsValidName(name string) (bool, error) {
	m, err := regexp.MatchString("[a-zA-Z0-9]{3,12}", name)
	if err != nil {
		return false, err
	}
	if !m {
		return false, nil
	}

	_, err = GetPropertiesByName(name)
	if err == nil {
		return false, nil
	}

	if err.Error() != "unable to find character by name" {
		return false, nil
	}

	bn, err := blocked_name.IsBlockedName(name)
	if bn {
		return false, err
	}

	return true, nil
}

func GetForWorld(accountId uint32, worldId byte) ([]Model, error) {
	cs, err := requestPropertiesByAccountAndWorld(accountId, worldId)
	if err != nil {
		return nil, err
	}

	var characters = make([]Model, 0)
	for _, x := range cs.DataList() {
		c, err := getFromProperties(&x)
		if err != nil {
			return nil, err
		}
		characters = append(characters, *c)
	}
	return characters, nil
}

func GetById(characterId uint32) (*Model, error) {
	cs, err := requestPropertiesById(characterId)
	if err != nil {
		return nil, err
	}

	c, err := getFromProperties(cs.Data())
	if err != nil {
		return nil, err
	}
	return c, nil
}

func getFromProperties(data *propertiesDataBody) (*Model, error) {
	ca := makeProperties(data)
	if ca == nil {
		return nil, errors.New("unable to make character properties")
	}

	eq, err := inventory.GetEquippedItemsForCharacter(ca.Id())
	if err != nil {
		return nil, err
	}

	ps, err := pet.GetForCharacter(nil)(ca.Id())
	if err != nil {
		return nil, err
	}

	c := NewCharacter(*ca, eq, ps)
	return &c, nil
}

func SeedCharacter(accountId uint32, worldId byte, name string, job uint32, face uint32, hair uint32, color uint32, skinColor uint32, gender byte, top uint32, bottom uint32, shoes uint32, weapon uint32) (*Properties, error) {
	ca, err := seedCharacter(accountId, worldId, name, job, face, hair, color, skinColor, gender, top, bottom, shoes, weapon)
	if err != nil {
		return nil, err
	}
	return makeProperties(ca), nil
}
