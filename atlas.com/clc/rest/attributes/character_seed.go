package attributes

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
