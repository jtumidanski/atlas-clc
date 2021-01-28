package domain

type Character struct {
	attributes CharacterAttributes
	equipment  []EquippedItem
	pets       []Pet
}

func (c Character) Attributes() CharacterAttributes {
	return c.attributes
}

func (c Character) Pets() []Pet {
	return c.pets
}

func (c Character) Equipment() []EquippedItem {
	return c.equipment
}

func NewCharacter(attributes CharacterAttributes, equipment []EquippedItem, pets []Pet) *Character {
	return &Character{attributes, equipment, pets}
}
