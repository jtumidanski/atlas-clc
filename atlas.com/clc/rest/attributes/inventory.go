package attributes

import "strconv"

const (
	EquipmentAttributesType string = "com.atlas.cos.rest.attribute.EquipmentAttributes"
	EquipmentStatisticsType string = "com.atlas.cos.rest.attribute.EquipmentStatisticsAttributes"
)

var equipmentIncludes = []conditionalMapperProvider{
	transformEquipmentAttributes,
	transformEquipmentStatistics,
}

type InventoryDataContainer struct {
	data     dataSegment
	included dataSegment
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
	d, i, err := unmarshalRoot(data, mapperFunc(EmptyInventoryData), equipmentIncludes...)
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

func (c *InventoryDataContainer) GetIncludedEquippedItems() []EquipmentData {
	var e = make([]EquipmentData, 0)
	for _, x := range c.included {
		if val, ok := x.(*EquipmentData); ok && val.Attributes.Slot < 0 {
			e = append(e, *val)
		}
	}
	return e
}

func (c *InventoryDataContainer) GetEquipmentStatistics(id int) *EquipmentStatisticsAttributes {
	for _, x := range c.included {
		if val, ok := x.(*EquipmentStatisticsData); ok {
			eid, err := strconv.Atoi(val.Id)
			if err == nil && eid == id {
				return &val.Attributes
			}
		}
	}
	return nil
}

func transformEquipmentAttributes() (string, objectMapper) {
	return unmarshalData(EquipmentAttributesType, EmptyEquipmentData)
}

func transformEquipmentStatistics() (string, objectMapper) {
	return unmarshalData(EquipmentStatisticsType, EmptyEquipmentStatisticsData)
}

func EmptyInventoryData() interface{} {
	return &InventoryData{}
}
