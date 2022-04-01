package world

import (
	"atlas-clc/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"strconv"
)

func GetAll(l logrus.FieldLogger, span opentracing.Span) ([]Model, error) {
	return requests.SliceProvider[attributes, Model](l, span)(requestWorlds(), makeWorld)()
}

func GetById(l logrus.FieldLogger, span opentracing.Span) func(worldId byte) (Model, error) {
	return func(worldId byte) (Model, error) {
		return requests.Provider[attributes, Model](l, span)(requestWorld(worldId), makeWorld)()
	}
}

func makeWorld(data requests.DataBody[attributes]) (Model, error) {
	wid, err := strconv.Atoi(data.Id)
	if err != nil {
		return Model{}, err
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
	return w, nil
}

func GetCapacityStatus(l logrus.FieldLogger, span opentracing.Span) func(worldId byte) uint16 {
	return func(worldId byte) uint16 {
		w, err := GetById(l, span)(worldId)
		if err != nil {
			return StatusFull
		}
		return w.CapacityStatus()
	}
}
