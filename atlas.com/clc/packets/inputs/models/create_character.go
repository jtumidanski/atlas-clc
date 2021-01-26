package models

type CreateCharacter struct {
	name      string
	job       uint32
	face      uint32
	hair      uint32
	hairColor uint32
	skinColor uint32
	top       uint32
	bottom    uint32
	shoes     uint32
	weapon    uint32
	gender    uint32
}

func (c CreateCharacter) Name() string {
	return c.name
}

func (c CreateCharacter) Job() uint32 {
	return c.job
}

func (c CreateCharacter) Face() uint32 {
	return c.face
}

func (c CreateCharacter) Hair() uint32 {
	return c.hair
}

func (c CreateCharacter) HairColor() uint32 {
	return c.hairColor
}

func (c CreateCharacter) SkinColor() uint32 {
	return c.hairColor
}

func (c CreateCharacter) Gender() uint32 {
	return c.gender
}

func (c CreateCharacter) Top() uint32 {
	return c.top
}

func (c CreateCharacter) Bottom() uint32 {
	return c.bottom
}

func (c CreateCharacter) Shoes() uint32 {
	return c.shoes
}

func (c CreateCharacter) Weapon() uint32 {
	return c.weapon
}

func NewCreateCharacter(name string, job uint32, face uint32, hair uint32, hairColor uint32, skinColor uint32, top uint32, bottom uint32, shoes uint32, weapon uint32, gender uint32) *CreateCharacter {
	return &CreateCharacter{name, job, face, hair, hairColor, skinColor, top, bottom, shoes, weapon, gender}
}
