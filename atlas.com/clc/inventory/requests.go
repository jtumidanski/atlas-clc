package inventory

import (
	"atlas-clc/rest/requests"
	"fmt"
)

const (
	CharactersServicePrefix     string = "/ms/cos/"
	CharactersService                  = requests.BaseRequest + CharactersServicePrefix
	CharactersResource                 = CharactersService + "characters/"
	CharactersInventoryResource        = CharactersResource + "%d/inventories/"
	CharacterEquippedItems             = CharactersInventoryResource + "?type=equip&include=inventoryItems,equipmentStatistics"
)

func requestEquippedItemsForCharacter(characterId uint32) (*InventoryDataContainer, error) {
	ar := &InventoryDataContainer{}
	err := requests.Get(fmt.Sprintf(CharacterEquippedItems, characterId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
