package handler

import (
	"atlas-clc/account"
	"atlas-clc/character"
	"atlas-clc/session"
	"atlas-clc/socket/response/writer"
	"github.com/jtumidanski/atlas-socket/request"
	"github.com/sirupsen/logrus"
)

const OpCodeCharacterCreate uint16 = 0x16

type CharacterCreateRequest struct {
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
	gender    byte
}

func (c CharacterCreateRequest) Name() string {
	return c.name
}

func (c CharacterCreateRequest) Job() uint32 {
	return c.job
}

func (c CharacterCreateRequest) Face() uint32 {
	return c.face
}

func (c CharacterCreateRequest) Hair() uint32 {
	return c.hair
}

func (c CharacterCreateRequest) HairColor() uint32 {
	return c.hairColor
}

func (c CharacterCreateRequest) SkinColor() uint32 {
	return c.skinColor
}

func (c CharacterCreateRequest) Gender() byte {
	return c.gender
}

func (c CharacterCreateRequest) Top() uint32 {
	return c.top
}

func (c CharacterCreateRequest) Bottom() uint32 {
	return c.bottom
}

func (c CharacterCreateRequest) Shoes() uint32 {
	return c.shoes
}

func (c CharacterCreateRequest) Weapon() uint32 {
	return c.weapon
}

func ReadCharacterCreateRequest(reader *request.RequestReader) *CharacterCreateRequest {
	name := reader.ReadAsciiString()
	job := reader.ReadUint32()
	face := reader.ReadUint32()
	hair := reader.ReadUint32()
	hairColor := reader.ReadUint32()
	skinColor := reader.ReadUint32()
	top := reader.ReadUint32()
	bottom := reader.ReadUint32()
	shoes := reader.ReadUint32()
	weapon := reader.ReadUint32()
	gender := reader.ReadByte()
	return &CharacterCreateRequest{
		name,
		job,
		face,
		hair,
		hairColor,
		skinColor,
		top,
		bottom,
		shoes,
		weapon,
		gender,
	}
}

type CharacterCreateHandler struct {
}

func (h *CharacterCreateHandler) IsValid(l logrus.FieldLogger, ms *session.MapleSession) bool {
	v := account.IsLoggedIn((*ms).AccountId())
	if !v {
		l.Errorf("Attempting to process a [CharacterCreateRequest] when the account %d is not logged in.", (*ms).SessionId())
	}
	return v
}

func (h *CharacterCreateHandler) HandleRequest(l logrus.FieldLogger, ms *session.MapleSession, r *request.RequestReader) {
	p := ReadCharacterCreateRequest(r)

	ca, err := character.SeedCharacter((*ms).AccountId(), (*ms).WorldId(), p.Name(), p.Job(), p.Face(), p.Hair(), p.HairColor(), p.SkinColor(), p.Gender(), p.Top(), p.Bottom(), p.Shoes(), p.Weapon())
	if err != nil {
		l.WithError(err).Errorf("While seeding character")
		return
	}

	c, err := character.GetById(ca.Id())
	if err != nil {
		l.WithError(err).Errorf("Retrieving newly seeded character")
		return
	}

	(*ms).Announce(writer.WriteCharacterViewAddNew(*c))
}
