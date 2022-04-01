package character

import (
	"atlas-clc/character/inventory"
	"atlas-clc/character/properties"
	"atlas-clc/pet"
)

type Model struct {
	properties properties.Model
	equipment  []inventory.EquippedItem
	pets       []pet.Model
}

func (c Model) Properties() properties.Model {
	return c.properties
}

func (c Model) Pets() []pet.Model {
	return c.pets
}

func (c Model) Equipment() []inventory.EquippedItem {
	return c.equipment
}

func NewCharacter(attributes properties.Model, equipment []inventory.EquippedItem, pets []pet.Model) Model {
	return Model{attributes, equipment, pets}
}
