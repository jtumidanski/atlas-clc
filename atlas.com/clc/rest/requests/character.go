package requests

import (
	"atlas-clc/rest/attributes"
	"fmt"
	"log"
)

func GetCharacterAttributesByName(l *log.Logger, name string) (*attributes.CharacterAttributesDataContainer, error) {
	ar := &attributes.CharacterAttributesDataContainer{}
	err := Get(l, fmt.Sprintf("http://atlas-nginx:80/ms/cos/characters?name=%s", name), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func GetCharacterAttributesForAccountByWorld(l *log.Logger, accountId int, worldId byte) (*attributes.CharacterAttributesDataContainer, error) {
	ar := &attributes.CharacterAttributesDataContainer{}
	err := Get(l, fmt.Sprintf("http://atlas-nginx:80/ms/cos/characters?accountId=%d&worldId=%d", accountId, worldId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func GetCharacterAttributesById(l *log.Logger, characterId uint32) (*attributes.CharacterAttributesDataContainer, error) {
	ar := &attributes.CharacterAttributesDataContainer{}
	err := Get(l, fmt.Sprintf("http://atlas-nginx:80/ms/cos/characters/%d", characterId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func GetEquippedItemsForCharacter(l *log.Logger, characterId uint32) (*attributes.InventoryDataContainer, error) {
	ar := &attributes.InventoryDataContainer{}
	err := Get(l, fmt.Sprintf("http://atlas-nginx:80/ms/cos/characters/%d/inventories?type=equip&include=inventoryItems,equipmentStatistics", characterId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func SeedCharacter(l *log.Logger, accountId int, worldId byte, name string, job uint32, face uint32, hair uint32, color uint32, skinColor uint32, gender byte, top uint32, bottom uint32, shoes uint32, weapon uint32) (*attributes.CharacterAttributesData, error) {
	i := attributes.CharacterSeedAttributesInputDataContainer{
		Data: attributes.CharacterSeedAttributesData{
			Id:   "0",
			Type: "com.atlas.cos.rest.attribute.CharacterSeedAttributes",
			Attributes: attributes.CharacterSeedAttributesAttributes{
				AccountId: accountId,
				WorldId:   worldId,
				Name:      name,
				JobIndex:  job,
				Face:      face,
				Hair:      hair,
				HairColor: color,
				Skin:      skinColor,
				Gender:    gender,
				Top:       top,
				Bottom:    bottom,
				Shoes:     shoes,
				Weapon:    weapon,
			},
		},
	}

	r, err := Post(l, "http://atlas-nginx:80/ms/cos/characters/seeds", i)
	if err != nil {
		return nil, err
	}

	ca := &attributes.CharacterAttributesDataContainer{}
	err = ProcessResponse(l, r, ca)
	if err != nil {
		return nil, err
	}
	return ca.Data(), nil
}
