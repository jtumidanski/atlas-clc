package inventory

import (
	"atlas-clc/rest/requests"
	"fmt"
)

const (
	charactersServicePrefix     string = "/ms/cos/"
	charactersService                  = requests.BaseRequest + charactersServicePrefix
	charactersResource                 = charactersService + "characters/"
	charactersInventoryResource        = charactersResource + "%d/inventories/"
	characterItems                     = charactersInventoryResource + "?type=%s&include=inventoryItems,equipmentStatistics"
)

func requestEquippedItemsForCharacter(characterId uint32) requests.Request[inventoryAttributes] {
	return requestItemsForCharacter(characterId, "equip")
}

func requestItemsForCharacter(characterId uint32, inventoryType string) requests.Request[inventoryAttributes] {
	return requests.MakeGetRequest[inventoryAttributes](fmt.Sprintf(characterItems, characterId, inventoryType), requests.AddMappers(equipmentIncludes))
}
