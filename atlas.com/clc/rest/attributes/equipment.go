package attributes

type EquipmentListDataContainer struct {
	Data []EquipmentData `json:"data"`
}

type EquipmentDataContainer struct {
	Data EquipmentData `json:"data"`
}

type EquipmentInputDataContainer struct {
	Data EquipmentData `json:"data"`
}

type EquipmentData struct {
	Id         string              `json:"id"`
	Type       string              `json:"type"`
	Attributes EquipmentAttributes `json:"attributes"`
}

type EquipmentAttributes struct {
	EquipmentId int   `json:"equipmentId"`
	Slot        int16 `json:"slot"`
}

type EquipmentStatisticsListDataContainer struct {
	Data []EquipmentStatisticsData `json:"data"`
}

type EquipmentStatisticsDataContainer struct {
	Data EquipmentStatisticsData `json:"data"`
}

type EquipmentStatisticsInputDataContainer struct {
	Data EquipmentStatisticsData `json:"data"`
}

type EquipmentStatisticsData struct {
	Id         string                        `json:"id"`
	Type       string                        `json:"type"`
	Attributes EquipmentStatisticsAttributes `json:"attributes"`
}

type EquipmentStatisticsAttributes struct {
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
