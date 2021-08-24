package character

import (
	"atlas-clc/character/properties"
	"atlas-clc/rest/requests"
)

const (
	CharactersServicePrefix string = "/ms/cos/"
	CharactersService              = requests.BaseRequest + CharactersServicePrefix
	CharactersResource             = CharactersService + "characters/"
	CharacterSeeds                 = CharactersResource + "seeds"
)

func seedCharacter(accountId uint32, worldId byte, name string, job uint32, face uint32, hair uint32, color uint32, skinColor uint32, gender byte, top uint32, bottom uint32, shoes uint32, weapon uint32) (*properties.DataBody, error) {
	i := seedInputDataContainer{
		Data: seedDataBody{
			Id:   "0",
			Type: "com.atlas.cos.rest.attribute.CharacterSeedAttributes",
			Attributes: seedAttributes{
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

	ca := &properties.DataContainer{}
	err = requests.ProcessResponse(r, ca)
	if err != nil {
		return nil, err
	}
	return ca.Data(), nil
}
