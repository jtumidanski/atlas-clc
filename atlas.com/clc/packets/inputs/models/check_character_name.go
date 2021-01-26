package models

type CheckCharacterName struct {
	name string
}

func (c *CheckCharacterName) Name() string {
	return c.name
}

func NewCheckCharacterName(name string) *CheckCharacterName {
	return &CheckCharacterName{name}
}
