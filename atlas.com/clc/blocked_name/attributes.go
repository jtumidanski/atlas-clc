package blocked_name

import "atlas-clc/rest/response"

type BlockedNameDataContainer struct {
	data response.DataSegment
}

type BlockedNameData struct {
	Id         string                `json:"id"`
	Type       string                `json:"type"`
	Attributes BlockedNameAttributes `json:"attributes"`
}

type BlockedNameAttributes struct {
	Name string `json:"name"`
}

func (b *BlockedNameDataContainer) UnmarshalJSON(data []byte) error {
	d, _, err := response.UnmarshalRoot(data, response.MapperFunc(EmptyBlockedNameData))
	if err != nil {
		return err
	}

	b.data = d
	return nil
}

func (b *BlockedNameDataContainer) Data() *BlockedNameData {
	if len(b.data) >= 1 {
		return b.data[0].(*BlockedNameData)
	}
	return nil
}

func (b *BlockedNameDataContainer) DataList() []BlockedNameData {
	var r = make([]BlockedNameData, 0)
	for _, x := range b.data {
		r = append(r, *x.(*BlockedNameData))
	}
	return r
}

func EmptyBlockedNameData() interface{} {
	return &BlockedNameData{}
}
