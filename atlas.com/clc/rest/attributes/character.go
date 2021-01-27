package attributes

import "encoding/json"

type CharacterAttributesDataContainer struct {
	data []CharacterAttributesData
}

func (c *CharacterAttributesDataContainer) UnmarshalJSON(data []byte) error {
	c.data = make([]CharacterAttributesData, 0)

	var single = struct{ Data CharacterAttributesData }{}
	err := json.Unmarshal(data, &single)
	if err == nil {
		c.data = append(c.data, single.Data)
		return nil
	} else {
		var list = struct{ Data []CharacterAttributesData }{}
		err = json.Unmarshal(data, &list)
		if err == nil {
			c.data = list.Data
			return nil
		}
		return err
	}
}

func (c *CharacterAttributesDataContainer) Data() *CharacterAttributesData {
	if len(c.data) >= 1 {
		return &c.data[0]
	}
	return nil
}

func (c *CharacterAttributesDataContainer) DataList() []CharacterAttributesData {
	return c.data
}

type CharacterAttributesInputDataContainer struct {
	Data CharacterAttributesData `json:"data"`
}

type CharacterAttributesData struct {
	Id         string                        `json:"id"`
	Type       string                        `json:"type"`
	Attributes CharacterAttributesAttributes `json:"attributes"`
}

type CharacterAttributesAttributes struct {
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
	Meso               int    `json:"meso"`
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

type CharacterSeedAttributesListDataContainer struct {
	Data []CharacterSeedAttributesData `json:"data"`
}

type CharacterSeedAttributesDataContainer struct {
	Data CharacterSeedAttributesData `json:"data"`
}

type CharacterSeedAttributesInputDataContainer struct {
	Data CharacterSeedAttributesData `json:"data"`
}

type CharacterSeedAttributesData struct {
	Id         string                            `json:"id"`
	Type       string                            `json:"type"`
	Attributes CharacterSeedAttributesAttributes `json:"attributes"`
}

type CharacterSeedAttributesAttributes struct {
	AccountId int    `json:"accountId"`
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
