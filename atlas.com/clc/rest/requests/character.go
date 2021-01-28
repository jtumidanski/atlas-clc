package requests

import (
	"atlas-clc/rest/attributes"
	"fmt"
)

const (
	CharactersServicePrefix     string = "/ms/cos/"
	CharactersService                  = BaseRequest + CharactersServicePrefix
	CharactersResource                 = CharactersService + "characters/"
	CharactersByName                   = CharactersResource + "?name=%s"
	CharactersForAccountByWorld        = CharactersResource + "?accountId=%d&worldId=%d"
	CharactersById                     = CharactersResource + "%d"
	CharactersInventoryResource        = CharactersResource + "%d/inventories/"
	CharacterEquippedItems             = CharactersInventoryResource + "?type=equip&include=inventoryItems,equipmentStatistics"
	CharacterSeeds                     = CharactersResource + "seeds/"
)

func GetCharacterAttributesByName(name string) (*attributes.CharacterAttributesDataContainer, error) {
	ar := &attributes.CharacterAttributesDataContainer{}
	err := Get(fmt.Sprintf(CharactersByName, name), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func GetCharacterAttributesForAccountByWorld(accountId int, worldId byte) (*attributes.CharacterAttributesDataContainer, error) {
	ar := &attributes.CharacterAttributesDataContainer{}
	err := Get(fmt.Sprintf(CharactersForAccountByWorld, accountId, worldId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func GetCharacterAttributesById(characterId uint32) (*attributes.CharacterAttributesDataContainer, error) {
	ar := &attributes.CharacterAttributesDataContainer{}
	err := Get(fmt.Sprintf(CharactersById, characterId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func GetEquippedItemsForCharacter(characterId uint32) (*attributes.InventoryDataContainer, error) {
	ar := &attributes.InventoryDataContainer{}
	err := Get(fmt.Sprintf(CharacterEquippedItems, characterId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func SeedCharacter(accountId int, worldId byte, name string, job uint32, face uint32, hair uint32, color uint32, skinColor uint32, gender byte, top uint32, bottom uint32, shoes uint32, weapon uint32) (*attributes.CharacterAttributesData, error) {
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

	r, err := Post(CharacterSeeds, i)
	if err != nil {
		return nil, err
	}

	ca := &attributes.CharacterAttributesDataContainer{}
	err = ProcessResponse(r, ca)
	if err != nil {
		return nil, err
	}
	return ca.Data(), nil
}
