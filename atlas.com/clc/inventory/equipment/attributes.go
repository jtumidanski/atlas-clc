package equipment

type DataBody struct {
	Id         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	EquipmentId int   `json:"equipmentId"`
	Slot        int16 `json:"slot"`
}

func EmptyEquipmentData() interface{} {
	return &DataBody{}
}

type StatisticsDataBody struct {
	Id         string               `json:"id"`
	Type       string               `json:"type"`
	Attributes StatisticsAttributes `json:"attributes"`
}

type StatisticsAttributes struct {
	ItemId        uint32 `json:"itemId"`
	Strength      uint16 `json:"strength"`
	Dexterity     uint16 `json:"dexterity"`
	Intelligence  uint16 `json:"intelligence"`
	Luck          uint16 `json:"luck"`
	Hp            uint16 `json:"hp"`
	Mp            uint16 `json:"mp"`
	WeaponAttack  uint16 `json:"weaponAttack"`
	MagicAttack   uint16 `json:"magicAttack"`
	WeaponDefense uint16 `json:"weaponDefense"`
	MagicDefense  uint16 `json:"magicDefense"`
	Accuracy      uint16 `json:"accuracy"`
	Avoidability  uint16 `json:"avoidability"`
	Hands         uint16 `json:"hands"`
	Speed         uint16 `json:"speed"`
	Jump          uint16 `json:"jump"`
	Slots         uint16 `json:"slots"`
}

func EmptyEquipmentStatisticsData() interface{} {
	return &StatisticsDataBody{}
}
