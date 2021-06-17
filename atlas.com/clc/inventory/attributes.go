package inventory

import (
	"atlas-clc/inventory/equipment"
	"atlas-clc/rest/response"
	"strconv"
)

const (
	EquipmentAttributesType string = "com.atlas.cos.rest.attribute.EquipmentAttributes"
	EquipmentStatisticsType string = "com.atlas.cos.rest.attribute.EquipmentStatisticsAttributes"
)

var equipmentIncludes = []response.ConditionalMapperProvider{
	transformEquipmentAttributes,
	transformEquipmentStatistics,
}

type InventoryDataContainer struct {
	data     response.DataSegment
	included response.DataSegment
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

func (c *InventoryDataContainer) UnmarshalJSON(data []byte) error {
	d, i, err := response.UnmarshalRoot(data, response.MapperFunc(EmptyInventoryData), equipmentIncludes...)
	if err != nil {
		return err
	}

	c.data = d
	c.included = i
	return nil
}

func (c *InventoryDataContainer) Data() *InventoryData {
	if len(c.data) >= 1 {
		return c.data[0].(*InventoryData)
	}
	return nil
}

func (c *InventoryDataContainer) DataList() []InventoryData {
	var r = make([]InventoryData, 0)
	for _, x := range c.data {
		r = append(r, *x.(*InventoryData))
	}
	return r
}

func (c *InventoryDataContainer) GetIncludedEquippedItems() []equipment.EquipmentData {
	var e = make([]equipment.EquipmentData, 0)
	for _, x := range c.included {
		if val, ok := x.(*equipment.EquipmentData); ok && val.Attributes.Slot < 0 {
			e = append(e, *val)
		}
	}
	return e
}

func (c *InventoryDataContainer) GetEquipmentStatistics(id int) *equipment.EquipmentStatisticsAttributes {
	for _, x := range c.included {
		if val, ok := x.(*equipment.EquipmentStatisticsData); ok {
			eid, err := strconv.Atoi(val.Id)
			if err == nil && eid == id {
				return &val.Attributes
			}
		}
	}
	return nil
}

func transformEquipmentAttributes() (string, response.ObjectMapper) {
	return response.UnmarshalData(EquipmentAttributesType, equipment.EmptyEquipmentData)
}

func transformEquipmentStatistics() (string, response.ObjectMapper) {
	return response.UnmarshalData(EquipmentStatisticsType, equipment.EmptyEquipmentStatisticsData)
}

func EmptyInventoryData() interface{} {
	return &InventoryData{}
}
