package inventory

import (
	"atlas-clc/inventory/equipment"
	"atlas-clc/rest/response"
	"encoding/json"
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

type dataContainer struct {
	data     response.DataSegment
	included response.DataSegment
}

type dataBody struct {
	Id         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes attributes `json:"attributes"`
}

type attributes struct {
	Type     string `json:"type"`
	Capacity int    `json:"capacity"`
}

func (c *dataContainer) MarshalJSON() ([]byte, error) {
	t := struct {
		Data     interface{} `json:"data"`
		Included interface{} `json:"included"`
	}{}
	if len(c.data) == 1 {
		t.Data = c.data[0]
	} else {
		t.Data = c.data
	}
	return json.Marshal(t)
}

func (c *dataContainer) UnmarshalJSON(data []byte) error {
	d, i, err := response.UnmarshalRoot(data, response.MapperFunc(EmptyInventoryData), equipmentIncludes...)
	if err != nil {
		return err
	}

	c.data = d
	c.included = i
	return nil
}

func (c *dataContainer) Data() *dataBody {
	if len(c.data) >= 1 {
		return c.data[0].(*dataBody)
	}
	return nil
}

func (c *dataContainer) DataList() []dataBody {
	var r = make([]dataBody, 0)
	for _, x := range c.data {
		r = append(r, *x.(*dataBody))
	}
	return r
}

func (c *dataContainer) GetIncludedEquippedItems() []equipment.DataBody {
	var e = make([]equipment.DataBody, 0)
	for _, x := range c.included {
		if val, ok := x.(*equipment.DataBody); ok && val.Attributes.Slot < 0 {
			e = append(e, *val)
		}
	}
	return e
}

func (c *dataContainer) GetEquipmentStatistics(id int) *equipment.StatisticsAttributes {
	for _, x := range c.included {
		if val, ok := x.(*equipment.StatisticsDataBody); ok {
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
	return &dataBody{}
}
