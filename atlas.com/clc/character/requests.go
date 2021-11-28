package character

import (
	"atlas-clc/character/properties"
	"atlas-clc/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	CharactersServicePrefix string = "/ms/cos/"
	CharactersService              = requests.BaseRequest + CharactersServicePrefix
	CharactersResource             = CharactersService + "characters/"
	CharacterSeeds                 = CharactersResource + "seeds"
)

func seedCharacter(l logrus.FieldLogger, span opentracing.Span) func(accountId uint32, worldId byte, name string, job uint32, face uint32, hair uint32, color uint32, skinColor uint32, gender byte, top uint32, bottom uint32, shoes uint32, weapon uint32) (*properties.DataBody, error) {
	return func(accountId uint32, worldId byte, name string, job uint32, face uint32, hair uint32, color uint32, skinColor uint32, gender byte, top uint32, bottom uint32, shoes uint32, weapon uint32) (*properties.DataBody, error) {
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

		resp := &properties.DataContainer{}
		errResp := &requests.ErrorListDataContainer{}

		err := requests.Post(l, span)(CharacterSeeds, i, resp, errResp)
		if err != nil {
			return nil, err
		}
		return resp.Data(), nil
	}
}
