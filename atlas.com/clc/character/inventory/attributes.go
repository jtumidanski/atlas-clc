package inventory

import (
	"atlas-clc/rest/requests"
	"atlas-clc/rest/response"
)

const (
	EquipmentAttributesType string = "com.atlas.cos.rest.attribute.EquipmentAttributes"
	EquipmentStatisticsType string = "com.atlas.cos.rest.attribute.EquipmentStatisticsAttributes"
)

var equipmentIncludes = []response.ConditionalMapperProvider{
	transformEquipmentAttributes,
	transformEquipmentStatistics,
}

func transformEquipmentAttributes() (string, response.ObjectMapper) {
	return response.UnmarshalData(EquipmentAttributesType, func() interface{} {
		return &requests.DataBody[equipmentAttributes]{}
	})
}

func transformEquipmentStatistics() (string, response.ObjectMapper) {
	return response.UnmarshalData(EquipmentStatisticsType, func() interface{} {
		return &requests.DataBody[equipmentStatisticsAttributes]{}
	})
}

type inventoryAttributes struct {
	Type     string `json:"type"`
	Capacity byte   `json:"capacity"`
}

type equipmentAttributes struct {
	EquipmentId int   `json:"equipmentId"`
	Slot        int16 `json:"slot"`
}

type equipmentStatisticsAttributes struct {
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
	Slots         byte   `json:"slots"`
}
