package world

import "atlas-clc/rest/response"

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
	Name               string `json:"name"`
	Flag               int    `json:"flag"`
	Message            string `json:"message"`
	EventMessage       string `json:"eventMessage"`
	Recommended        bool   `json:"recommended"`
	RecommendedMessage string `json:"recommendedMessage"`
	CapacityStatus     uint16 `json:"capacityStatus"`
}

func (c *dataContainer) UnmarshalJSON(data []byte) error {
	d, i, err := response.UnmarshalRoot(data, response.MapperFunc(EmptyWorldData))
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

func EmptyWorldData() interface{} {
	return &dataBody{}
}