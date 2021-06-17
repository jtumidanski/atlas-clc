package blocked_name

import "atlas-clc/rest/response"

type dataContainer struct {
	data response.DataSegment
}

type dataBody struct {
	Id         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes attributes `json:"attributes"`
}

type attributes struct {
	Name string `json:"name"`
}

func (b *dataContainer) UnmarshalJSON(data []byte) error {
	d, _, err := response.UnmarshalRoot(data, response.MapperFunc(EmptyBlockedNameData))
	if err != nil {
		return err
	}

	b.data = d
	return nil
}

func (b *dataContainer) Data() *dataBody {
	if len(b.data) >= 1 {
		return b.data[0].(*dataBody)
	}
	return nil
}

func (b *dataContainer) DataList() []dataBody {
	var r = make([]dataBody, 0)
	for _, x := range b.data {
		r = append(r, *x.(*dataBody))
	}
	return r
}

func EmptyBlockedNameData() interface{} {
	return &dataBody{}
}
