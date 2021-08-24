package character

import (
	"atlas-clc/blocked_name"
	"atlas-clc/character/properties"
	"atlas-clc/inventory"
	"atlas-clc/pet"
	"github.com/sirupsen/logrus"
	"regexp"
)

func IsValidName(l logrus.FieldLogger) func(name string) (bool, error) {
	return func(name string) (bool, error) {
		m, err := regexp.MatchString("[a-zA-Z0-9]{3,12}", name)
		if err != nil {
			return false, err
		}
		if !m {
			return false, nil
		}

		_, err = properties.GetByName(l)(name)
		if err == nil {
			return false, nil
		}

		if err.Error() != "unable to find character by name" {
			return false, nil
		}

		bn, err := blocked_name.IsBlockedName(l)(name)
		if bn {
			return false, err
		}

		return true, nil
	}
}

func GetForWorld(l logrus.FieldLogger) func(accountId uint32, worldId byte) ([]Model, error) {
	return func(accountId uint32, worldId byte) ([]Model, error) {
		var characters = make([]Model, 0)
		cs, err := properties.GetForWorld(l)(accountId, worldId)
		if err != nil {
			return nil, err
		}
		for _, x := range cs {
			c, err := fromProperties(l)(&x)
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
		cs, err := properties.GetById(l)(characterId)
		if err != nil {
			return nil, err
		}

		c, err := fromProperties(l)(cs)
		if err != nil {
			return nil, err
		}
		return c, nil
	}
}

func fromProperties(l logrus.FieldLogger) func(data *properties.Model) (*Model, error) {
	return func(data *properties.Model) (*Model, error) {
		eq, err := inventory.GetEquippedItemsForCharacter(l)(data.Id())
		if err != nil {
			return nil, err
		}

		ps, err := pet.GetForCharacter(nil)(data.Id())
		if err != nil {
			return nil, err
		}

		c := NewCharacter(*data, eq, ps)
		return &c, nil
	}
}

func SeedCharacter(accountId uint32, worldId byte, name string, job uint32, face uint32, hair uint32, color uint32, skinColor uint32, gender byte, top uint32, bottom uint32, shoes uint32, weapon uint32) (*properties.Model, error) {
	ca, err := seedCharacter(accountId, worldId, name, job, face, hair, color, skinColor, gender, top, bottom, shoes, weapon)
	if err != nil {
		return nil, err
	}
	p, err := properties.MakeModel(ca)
	if err != nil {
		return nil, err
	}
	return p, nil
}
