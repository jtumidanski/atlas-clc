package character

import "atlas-clc/rest/response"

type propertiesDataContainer struct {
	data response.DataSegment
}

type propertiesDataBody struct {
	Id         string               `json:"id"`
	Type       string               `json:"type"`
	Attributes propertiesAttributes `json:"attributes"`
}

type propertiesAttributes struct {
	AccountId          int    `json:"accountId"`
	WorldId            byte   `json:"worldId"`
	Name               string `json:"name"`
	Level              byte   `json:"level"`
	Experience         uint32 `json:"experience"`
	GachaponExperience uint32 `json:"gachaponExperience"`
	Strength           uint16 `json:"strength"`
	Dexterity          uint16 `json:"dexterity"`
	Intelligence       uint16 `json:"intelligence"`
	Luck               uint16 `json:"luck"`
	Hp                 uint16 `json:"hp"`
	MaxHp              uint16 `json:"maxHp"`
	Mp                 uint16 `json:"mp"`
	MaxMp              uint16 `json:"maxMp"`
	Meso               uint32 `json:"meso"`
	HpMpUsed           int    `json:"hpMpUsed"`
	JobId              uint16 `json:"jobId"`
	SkinColor          byte   `json:"skinColor"`
	Gender             byte   `json:"gender"`
	Fame               int16  `json:"fame"`
	Hair               uint32 `json:"hair"`
	Face               uint32 `json:"face"`
	Ap                 uint16 `json:"ap"`
	Sp                 string `json:"sp"`
	MapId              uint32 `json:"mapId"`
	SpawnPoint         byte   `json:"spawnPoint"`
	Gm                 int    `json:"gm"`
	X                  int    `json:"x"`
	Y                  int    `json:"y"`
	Stance             byte   `json:"stance"`
}

func (c *propertiesDataContainer) UnmarshalJSON(data []byte) error {
	d, _, err := response.UnmarshalRoot(data, response.MapperFunc(EmptyCharacterAttributesData))
	if err != nil {
		return err
	}
	c.data = d
	return nil
}

func (c *propertiesDataContainer) Data() *propertiesDataBody {
	if len(c.data) >= 1 {
		return c.data[0].(*propertiesDataBody)
	}
	return nil
}

func (c *propertiesDataContainer) DataList() []propertiesDataBody {
	var r = make([]propertiesDataBody, 0)
	for _, x := range c.data {
		r = append(r, *x.(*propertiesDataBody))
	}
	return r
}

func EmptyCharacterAttributesData() interface{} {
	return &propertiesDataBody{}
}

type seedInputDataContainer struct {
	Data seedDataBody `json:"data"`
}

type seedDataBody struct {
	Id         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes seedAttributes `json:"properties"`
}

type seedAttributes struct {
	AccountId uint32 `json:"accountId"`
	WorldId   byte   `json:"worldId"`
	Name      string `json:"name"`
	JobIndex  uint32 `json:"jobIndex"`
	Face      uint32 `json:"face"`
	Hair      uint32 `json:"hair"`
	HairColor uint32 `json:"hairColor"`
	Skin      uint32 `json:"skin"`
	Gender    byte   `json:"gender"`
	Top       uint32 `json:"top"`
	Bottom    uint32 `json:"bottom"`
	Shoes     uint32 `json:"shoes"`
	Weapon    uint32 `json:"weapon"`
}
