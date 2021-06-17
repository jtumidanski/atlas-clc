package character

import (
	"atlas-clc/rest/requests"
	"fmt"
)

const (
	CharactersServicePrefix     string = "/ms/cos/"
	CharactersService                  = requests.BaseRequest + CharactersServicePrefix
	CharactersResource                 = CharactersService + "characters/"
	CharactersByName                   = CharactersResource + "?name=%s"
	CharactersForAccountByWorld        = CharactersResource + "?accountId=%d&worldId=%d"
	CharactersById                     = CharactersResource + "%d"
	CharactersInventoryResource        = CharactersResource + "%d/inventories/"
	CharacterEquippedItems             = CharactersInventoryResource + "?type=equip&include=inventoryItems,equipmentStatistics"
	CharacterSeeds                     = CharactersResource + "seeds"
)

func requestCharacterAttributesByName(name string) (*CharacterAttributesDataContainer, error) {
	ar := &CharacterAttributesDataContainer{}
	err := requests.Get(fmt.Sprintf(CharactersByName, name), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func requestCharacterAttributesForAccountByWorld(accountId uint32, worldId byte) (*CharacterAttributesDataContainer, error) {
	ar := &CharacterAttributesDataContainer{}
	err := requests.Get(fmt.Sprintf(CharactersForAccountByWorld, accountId, worldId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func requestCharacterAttributesById(characterId uint32) (*CharacterAttributesDataContainer, error) {
	ar := &CharacterAttributesDataContainer{}
	err := requests.Get(fmt.Sprintf(CharactersById, characterId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func seedCharacter(accountId uint32, worldId byte, name string, job uint32, face uint32, hair uint32, color uint32, skinColor uint32, gender byte, top uint32, bottom uint32, shoes uint32, weapon uint32) (*CharacterAttributesData, error) {
	i := CharacterSeedAttributesInputDataContainer{
		Data: CharacterSeedAttributesData{
			Id:   "0",
			Type: "com.atlas.cos.rest.attribute.CharacterSeedAttributes",
			Attributes: CharacterSeedAttributesAttributes{
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

	r, err := requests.Post(CharacterSeeds, i)
	if err != nil {
		return nil, err
	}

	ca := &CharacterAttributesDataContainer{}
	err = requests.ProcessResponse(r, ca)
	if err != nil {
		return nil, err
	}
	return ca.Data(), nil
}
