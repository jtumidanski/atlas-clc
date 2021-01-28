package processors

import (
	"atlas-clc/domain"
	"atlas-clc/rest/attributes"
	"atlas-clc/rest/requests"
	"strconv"
)

func GetWorlds() ([]domain.World, error) {
	r, err := requests.GetWorlds()
	if err != nil {
		return nil, err
	}

	var ws = make([]domain.World, 0)
	for _, x := range r.DataList() {
		w, err := makeWorld(x)
		if err == nil {
			ws = append(ws, *w)
		}
	}
	return ws, nil
}

func GetWorld(worldId byte) (*domain.World, error) {
	r, err := requests.GetWorld(worldId)
	if err != nil {
		return nil, err
	}

	return makeWorld(*r.Data())
}

func makeWorld(data attributes.WorldData) (*domain.World, error) {
	wid, err := strconv.Atoi(data.Id)
	if err != nil {
		return nil, err
	}

	att := data.Attributes
	return domain.NewWorldBuilder().
		SetId(byte(wid)).
		SetName(att.Name).
		SetFlag(att.Flag).
		SetMessage(att.Message).
		SetEventMessage(att.EventMessage).
		SetRecommended(att.Recommended).
		SetRecommendedMessage(att.RecommendedMessage).
		SetCapacityStatus(att.CapacityStatus).
		Build(), nil
}

func GetWorldCapacityStatus(worldId byte) uint16 {
	w, err := GetWorld(worldId)
	if err != nil {
		return domain.Full
	}
	return w.CapacityStatus()
}
