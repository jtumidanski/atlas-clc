package attributes

import (
	"encoding/json"
	"strconv"
)

type InventoryDataContainer struct {
	data     []InventoryData
	Included []interface{}
}

func (c *InventoryDataContainer) UnmarshalJSON(data []byte) error {
	c.data = make([]InventoryData, 0)

	var single = struct {
		Data     InventoryData
		Included []map[string]interface{}
	}{}
	err := json.Unmarshal(data, &single)
	if err == nil {
		c.data = append(c.data, single.Data)

		c.Included = make([]interface{}, 0)
		for _, x := range single.Included {
			c.Included = append(c.Included, addIncluded(x))
		}
		return nil
	} else {
		var list = struct {
			Data     []InventoryData
			Included []map[string]interface{}
		}{}
		err = json.Unmarshal(data, &list)
		if err == nil {
			c.data = list.Data
			c.Included = make([]interface{}, 0)
			for _, x := range list.Included {
				c.Included = append(c.Included, addIncluded(x))
			}
			return nil
		}
		return err
	}
}

func (c *InventoryDataContainer) GetIncludedEquippedItems() []EquipmentData {
	var e = make([]EquipmentData, 0)
	for _, x := range c.Included {
		if val, ok := x.(*EquipmentData); ok && val.Attributes.Slot < 0 {
			e = append(e, *val)
		}
	}
	return e
}

func (c *InventoryDataContainer) GetEquipmentStatistics(id int) *EquipmentStatisticsAttributes {
	for _, x := range c.Included {
		if val, ok := x.(*EquipmentStatisticsData); ok {
			eid, err := strconv.Atoi(val.Id)
			if err == nil && eid == id {
				return &val.Attributes
			}
		}
	}
	return nil
}

func getAttribute(x map[string]interface{}, attrName string) float64 {
	return x["attributes"].(map[string]interface{})[attrName].(float64)
}

func addIncluded(x map[string]interface{}) interface{} {
	switch x["type"].(string) {
	case "com.atlas.cos.rest.attribute.EquipmentAttributes":
		return &EquipmentData{
			Id:   x["id"].(string),
			Type: x["type"].(string),
			Attributes: EquipmentAttributes{
				EquipmentId: int(getAttribute(x, "equipmentId")),
				Slot:        int16(getAttribute(x, "slot")),
			},
		}
	case "com.atlas.cos.rest.attribute.EquipmentStatisticsAttributes":
		return &EquipmentStatisticsData{
			Id:   x["id"].(string),
			Type: x["type"].(string),
			Attributes: EquipmentStatisticsAttributes{
				ItemId:        uint32(getAttribute(x, "itemId")),
				Strength:      uint16(getAttribute(x, "strength")),
				Dexterity:     uint16(getAttribute(x, "dexterity")),
				Intelligence:  uint16(getAttribute(x, "intelligence")),
				Luck:          uint16(getAttribute(x, "luck")),
				Hp:            uint16(getAttribute(x, "hp")),
				Mp:            uint16(getAttribute(x, "mp")),
				WeaponAttack:  uint16(getAttribute(x, "weaponAttack")),
				MagicAttack:   uint16(getAttribute(x, "magicAttack")),
				WeaponDefense: uint16(getAttribute(x, "weaponDefense")),
				MagicDefense:  uint16(getAttribute(x, "magicDefense")),
				Accuracy:      uint16(getAttribute(x, "accuracy")),
				Avoidability:  uint16(getAttribute(x, "avoidability")),
				Hands:         uint16(getAttribute(x, "hands")),
				Speed:         uint16(getAttribute(x, "speed")),
				Jump:          uint16(getAttribute(x, "jump")),
				Slots:         uint16(getAttribute(x, "slots")),
			},
		}
	default:
		return nil
	}
}

func (c *InventoryDataContainer) Data() *InventoryData {
	if len(c.data) >= 1 {
		return &c.data[0]
	}
	return nil
}

func (c *InventoryDataContainer) DataList() []InventoryData {
	return c.data
}

type InventoryInputDataContainer struct {
	Data InventoryData `json:"data"`
}

type InventoryData struct {
	Id         string              `json:"id"`
	Type       string              `json:"type"`
	Attributes InventoryAttributes `json:"attributes"`
}

type InventoryAttributes struct {
	Type     string `json:"type"`
	Capacity int    `json:"capacity"`
}
