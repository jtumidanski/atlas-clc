package inventory

import (
	"atlas-clc/rest/requests"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	CharactersServicePrefix     string = "/ms/cos/"
	CharactersService                  = requests.BaseRequest + CharactersServicePrefix
	CharactersResource                 = CharactersService + "characters/"
	CharactersInventoryResource        = CharactersResource + "%d/inventories/"
	CharacterEquippedItems             = CharactersInventoryResource + "?type=equip&include=inventoryItems,equipmentStatistics"
)

func requestEquippedItemsForCharacter(l logrus.FieldLogger) func(characterId uint32) (*dataContainer, error) {
	return func(characterId uint32) (*dataContainer, error) {
		ar := &dataContainer{}
		err := requests.Get(l)(fmt.Sprintf(CharacterEquippedItems, characterId), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}
