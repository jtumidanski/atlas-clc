package attributes

type WorldDataContainer struct {
	data     dataSegment
	included dataSegment
}

type WorldData struct {
	Id         string          `json:"id"`
	Type       string          `json:"type"`
	Attributes WorldAttributes `json:"attributes"`
}

type WorldAttributes struct {
	Name               string `json:"name"`
	Flag               int    `json:"flag"`
	Message            string `json:"message"`
	EventMessage       string `json:"eventMessage"`
	Recommended        bool   `json:"recommended"`
	RecommendedMessage string `json:"recommendedMessage"`
	CapacityStatus     uint16 `json:"capacityStatus"`
}

func (c *WorldDataContainer) UnmarshalJSON(data []byte) error {
	d, i, err := unmarshalRoot(data, mapperFunc(EmptyWorldData))
	if err != nil {
		return err
	}

	c.data = d
	c.included = i
	return nil
}

func (c *WorldDataContainer) Data() *WorldData {
	if len(c.data) >= 1 {
		return c.data[0].(*WorldData)
	}
	return nil
}

func (c *WorldDataContainer) DataList() []WorldData {
	var r = make([]WorldData, 0)
	for _, x := range c.data {
		r = append(r, *x.(*WorldData))
	}
	return r
}

func EmptyWorldData() interface{} {
	return &WorldData{}
}
