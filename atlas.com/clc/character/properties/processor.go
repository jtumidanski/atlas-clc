package properties

import (
	"errors"
	"github.com/sirupsen/logrus"
	"strconv"
)

func GetByName(l logrus.FieldLogger) func(name string) (*Model, error) {
	return func(name string) (*Model, error) {
		ca, err := requestPropertiesByName(name)(l)
		if err != nil {
			return nil, err
		}
		if len(ca.DataList()) <= 0 {
			return nil, errors.New("unable to find character by name")
		}

		return MakeModel(ca.Data())
	}
}

func GetForWorld(l logrus.FieldLogger) func(accountId uint32, worldId byte) ([]Model, error) {
	return func(accountId uint32, worldId byte) ([]Model, error) {
		cs, err := requestPropertiesByAccountAndWorld(accountId, worldId)(l)
		if err != nil {
			return nil, err
		}

		var characters = make([]Model, 0)
		for _, x := range cs.DataList() {
			c, err := MakeModel(&x)
			if err != nil {
				return nil, err
			}
			characters = append(characters, *c)
		}
		return characters, nil
	}
}

func GetById(l logrus.FieldLogger) func(characterId uint32) (*Model, error) {
	return func(characterId uint32) (*Model, error) {
		cs, err := requestPropertiesById(characterId)(l)
		if err != nil {
			return nil, err
		}
		return MakeModel(cs.Data())
	}
}

func MakeModel(ca *DataBody) (*Model, error) {
	cid, err := strconv.ParseUint(ca.Id, 10, 32)
	if err != nil {
		return nil, err
	}
	att := ca.Attributes
	r := NewBuilder().
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
	return &r, nil
}