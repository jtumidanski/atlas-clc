package properties

import (
	"atlas-clc/model"
	"atlas-clc/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"strconv"
)

func GetByName(l logrus.FieldLogger, span opentracing.Span) func(name string) (Model, error) {
	return func(name string) (Model, error) {
		return requests.Provider[Attributes, Model](l, span)(requestPropertiesByName(name), MakeModel)()
	}
}

func GetForWorld(l logrus.FieldLogger, span opentracing.Span) func(accountId uint32, worldId byte) ([]Model, error) {
	return func(accountId uint32, worldId byte) ([]Model, error) {
		cs, err := requestPropertiesByAccountAndWorld(accountId, worldId)(l, span)
		if err != nil {
			return nil, err
		}

		var characters = make([]Model, 0)
		for _, x := range cs.DataList() {
			c, err := MakeModel(x)
			if err != nil {
				return nil, err
			}
			characters = append(characters, c)
		}
		return characters, nil
	}
}

func ByIdModelProvider(l logrus.FieldLogger, span opentracing.Span) func(id uint32) model.Provider[Model] {
	return func(id uint32) model.Provider[Model] {
		return requests.Provider[Attributes, Model](l, span)(requestPropertiesById(id), MakeModel)
	}
}

func GetById(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) (Model, error) {
	return func(characterId uint32) (Model, error) {
		return ByIdModelProvider(l, span)(characterId)()
	}
}

func MakeModel(ca requests.DataBody[Attributes]) (Model, error) {
	cid, err := strconv.ParseUint(ca.Id, 10, 32)
	if err != nil {
		return Model{}, err
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
	return r, nil
}
