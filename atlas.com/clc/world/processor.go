package world

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

func GetAll(l logrus.FieldLogger) ([]Model, error) {
	r, err := requestWorlds(l)
	if err != nil {
		return nil, err
	}

	var ws = make([]Model, 0)
	for _, x := range r.DataList() {
		w, err := makeWorld(x)
		if err == nil {
			ws = append(ws, *w)
		}
	}
	return ws, nil
}

func GetById(l logrus.FieldLogger) func(worldId byte) (*Model, error) {
	return func(worldId byte) (*Model, error) {
		r, err := requestWorld(l)(worldId)
		if err != nil {
			return nil, err
		}

		return makeWorld(*r.Data())
	}
}

func makeWorld(data dataBody) (*Model, error) {
	wid, err := strconv.Atoi(data.Id)
	if err != nil {
		return nil, err
	}

	att := data.Attributes
	w := NewWorldBuilder().
		SetId(byte(wid)).
		SetName(att.Name).
		SetFlag(att.Flag).
		SetMessage(att.Message).
		SetEventMessage(att.EventMessage).
		SetRecommended(att.Recommended).
		SetRecommendedMessage(att.RecommendedMessage).
		SetCapacityStatus(att.CapacityStatus).
		Build()
	return &w, nil
}

func GetWorldCapacityStatus(l logrus.FieldLogger) func(worldId byte) uint16 {
	return func(worldId byte) uint16 {
		w, err := GetById(l)(worldId)
		if err != nil {
			return StatusFull
		}
		return w.CapacityStatus()
	}
}
