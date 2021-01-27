package processors

import (
	"atlas-clc/models"
	"atlas-clc/rest/attributes"
	"atlas-clc/rest/requests"
	"log"
	"strconv"
)

func GetWorlds(l *log.Logger) ([]models.World, error) {
	r, err := requests.GetWorlds(l)
	if err != nil {
		return nil, err
	}

	var ws = make([]models.World, 0)
	for _, x := range r.DataList() {
		w, err := makeWorld(x)
		if err == nil {
			ws = append(ws, *w)
		}
	}
	return ws, nil
}

func GetWorld(l *log.Logger, worldId byte) (*models.World, error) {
	r, err := requests.GetWorld(l, worldId)
	if err != nil {
		return nil, err
	}

	return makeWorld(*r.Data())
}

func makeWorld(data attributes.WorldData) (*models.World, error) {
	wid, err := strconv.Atoi(data.Id)
	if err != nil {
		return nil, err
	}

	att := data.Attributes
	return models.NewWorldBuilder().
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

func GetWorldCapacityStatus(l *log.Logger, worldId byte) uint16 {
	w, err := GetWorld(l, worldId)
	if err != nil {
		l.Println("[WARN] error deciding capacity status, will return full")
		return models.Full
	}
	return w.CapacityStatus()
}
