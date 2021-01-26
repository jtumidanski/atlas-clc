package attributes

import "encoding/json"

type InventoryDataContainer struct {
	data     []InventoryData
	Included []struct{}
}

func (c *InventoryDataContainer) UnmarshalJSON(data []byte) error {
	c.data = make([]InventoryData, 0)

	var single = struct {
		Data     InventoryData
		Included []struct{}
	}{}
	err := json.Unmarshal(data, &single)
	if err == nil {
		c.data = append(c.data, single.Data)
		c.Included = single.Included
		return nil
	} else {
		var list = struct {
			Data     []InventoryData
			Included []struct{}
		}{}
		err = json.Unmarshal(data, &list)
		if err == nil {
			c.data = list.Data
			c.Included = list.Included
			return nil
		}
		return err
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
