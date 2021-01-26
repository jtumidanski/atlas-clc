package models

import (
	"strconv"
	"strings"
)

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

type CharacterAttributes struct {
	id                 uint32
	worldId            byte
	name               string
	gender             byte
	skinColor          byte
	face               uint32
	hair               uint32
	level              byte
	jobId              uint16
	strength           uint16
	dexterity          uint16
	intelligence       uint16
	luck               uint16
	hp                 uint16
	maxHp              uint16
	mp                 uint16
	maxMp              uint16
	ap                 uint16
	sp                 string
	experience         uint32
	fame               int16
	gachaponExperience uint32
	mapId              uint32
	spawnPoint         byte
	gm                 bool
	gmJob              bool
	rank               int
	rankMove           int
	jobRank            int
	jobRankMove        int
}

func (a CharacterAttributes) Gm() bool {
	return a.gm
}

func (a CharacterAttributes) GmJob() bool {
	return a.gmJob
}

func (a CharacterAttributes) Rank() int {
	return a.rank
}

func (a CharacterAttributes) RankMove() int {
	return a.rankMove
}

func (a CharacterAttributes) JobRank() int {
	return a.jobRank
}

func (a CharacterAttributes) JobRankMove() int {
	return a.jobRankMove
}

func (a CharacterAttributes) Id() uint32 {
	return a.id
}

func (a CharacterAttributes) Name() string {
	return a.name
}

func (a CharacterAttributes) Gender() byte {
	return a.gender
}

func (a CharacterAttributes) SkinColor() byte {
	return a.skinColor
}

func (a CharacterAttributes) Face() uint32 {
	return a.face
}

func (a CharacterAttributes) Hair() uint32 {
	return a.hair
}

func (a CharacterAttributes) Level() byte {
	return a.level
}

func (a CharacterAttributes) JobId() uint16 {
	return a.jobId
}

func (a CharacterAttributes) Strength() uint16 {
	return a.strength
}

func (a CharacterAttributes) Dexterity() uint16 {
	return a.dexterity
}

func (a CharacterAttributes) Intelligence() uint16 {
	return a.intelligence
}

func (a CharacterAttributes) Luck() uint16 {
	return a.luck
}

func (a CharacterAttributes) Hp() uint16 {
	return a.hp
}

func (a CharacterAttributes) MaxHp() uint16 {
	return a.maxHp
}

func (a CharacterAttributes) Mp() uint16 {
	return a.mp
}

func (a CharacterAttributes) MaxMp() uint16 {
	return a.maxMp
}

func (a CharacterAttributes) Ap() uint16 {
	return a.ap
}

func (a CharacterAttributes) HasSPTable() bool {
	switch a.jobId {
	case 2001:
		return true
	case 2200:
		return true
	case 2210:
		return true
	case 2211:
		return true
	case 2212:
		return true
	case 2213:
		return true
	case 2214:
		return true
	case 2215:
		return true
	case 2216:
		return true
	case 2217:
		return true
	case 2218:
		return true
	default:
		return false
	}
}

func (a CharacterAttributes) Sp() []uint16 {
	s := strings.Split(a.sp, ",")
	var sps = make([]uint16, 0)
	for _, x := range s {
		sp, err := strconv.ParseUint(x, 10, 16)
		if err == nil {
			sps = append(sps, uint16(sp))
		}
	}
	return sps
}

func (a CharacterAttributes) RemainingSp() uint16 {
	return a.Sp()[a.skillBook()]

}

func (a CharacterAttributes) skillBook() uint16 {
	if a.jobId >= 2210 && a.jobId <= 2218 {
		return a.jobId - 2209
	}
	return 0
}

func (a CharacterAttributes) Experience() uint32 {
	return a.experience
}

func (a CharacterAttributes) Fame() int16 {
	return a.fame
}

func (a CharacterAttributes) GachaponExperience() uint32 {
	return a.gachaponExperience
}

func (a CharacterAttributes) SpawnPoint() byte {
	return a.spawnPoint
}

func (a CharacterAttributes) WorldId() byte {
	return a.worldId
}

func (a CharacterAttributes) MapId() uint32 {
	return a.mapId
}

func NewCharacterAttributes(id uint32, worldId byte, name string, gender byte, skinColor byte, face uint32, hair uint32, level byte, jobId uint16, strength uint16, dexterity uint16, intelligence uint16, luck uint16, hp uint16, maxHp uint16, mp uint16, maxMp uint16, ap uint16, sp string, experience uint32, fame int16, gachaponExperience uint32, mapId uint32, spawnPoint byte) *CharacterAttributes {
	return &CharacterAttributes{id, worldId, name, gender, skinColor, face, hair, level, jobId, strength, dexterity, intelligence, luck, hp, maxHp, mp, maxMp, ap, sp, experience, fame, gachaponExperience, mapId, spawnPoint, false, false, 0, 0, 0, 0}
}
