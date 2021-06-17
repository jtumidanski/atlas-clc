package character

import (
	"atlas-clc/inventory"
	"atlas-clc/pet"
	"strconv"
	"strings"
)

type Model struct {
	attributes Attributes
	equipment  []inventory.EquippedItem
	pets       []pet.Model
}

func (c Model) Attributes() Attributes {
	return c.attributes
}

func (c Model) Pets() []pet.Model {
	return c.pets
}

func (c Model) Equipment() []inventory.EquippedItem {
	return c.equipment
}

func NewCharacter(attributes Attributes, equipment []inventory.EquippedItem, pets []pet.Model) Model {
	return Model{attributes, equipment, pets}
}

type Attributes struct {
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

func (a Attributes) Gm() bool {
	return a.gm
}

func (a Attributes) GmJob() bool {
	return a.gmJob
}

func (a Attributes) Rank() int {
	return a.rank
}

func (a Attributes) RankMove() int {
	return a.rankMove
}

func (a Attributes) JobRank() int {
	return a.jobRank
}

func (a Attributes) JobRankMove() int {
	return a.jobRankMove
}

func (a Attributes) Id() uint32 {
	return a.id
}

func (a Attributes) Name() string {
	return a.name
}

func (a Attributes) Gender() byte {
	return a.gender
}

func (a Attributes) SkinColor() byte {
	return a.skinColor
}

func (a Attributes) Face() uint32 {
	return a.face
}

func (a Attributes) Hair() uint32 {
	return a.hair
}

func (a Attributes) Level() byte {
	return a.level
}

func (a Attributes) JobId() uint16 {
	return a.jobId
}

func (a Attributes) Strength() uint16 {
	return a.strength
}

func (a Attributes) Dexterity() uint16 {
	return a.dexterity
}

func (a Attributes) Intelligence() uint16 {
	return a.intelligence
}

func (a Attributes) Luck() uint16 {
	return a.luck
}

func (a Attributes) Hp() uint16 {
	return a.hp
}

func (a Attributes) MaxHp() uint16 {
	return a.maxHp
}

func (a Attributes) Mp() uint16 {
	return a.mp
}

func (a Attributes) MaxMp() uint16 {
	return a.maxMp
}

func (a Attributes) Ap() uint16 {
	return a.ap
}

func (a Attributes) HasSPTable() bool {
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

func (a Attributes) Sp() []uint16 {
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

func (a Attributes) RemainingSp() uint16 {
	return a.Sp()[a.skillBook()]

}

func (a Attributes) skillBook() uint16 {
	if a.jobId >= 2210 && a.jobId <= 2218 {
		return a.jobId - 2209
	}
	return 0
}

func (a Attributes) Experience() uint32 {
	return a.experience
}

func (a Attributes) Fame() int16 {
	return a.fame
}

func (a Attributes) GachaponExperience() uint32 {
	return a.gachaponExperience
}

func (a Attributes) SpawnPoint() byte {
	return a.spawnPoint
}

func (a Attributes) WorldId() byte {
	return a.worldId
}

func (a Attributes) MapId() uint32 {
	return a.mapId
}

type characterAttributeBuilder struct {
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

func NewCharacterAttributeBuilder() *characterAttributeBuilder {
	return &characterAttributeBuilder{}
}

func (c *characterAttributeBuilder) SetId(id uint32) *characterAttributeBuilder {
	c.id = id
	return c
}

func (c *characterAttributeBuilder) SetWorldId(worldId byte) *characterAttributeBuilder {
	c.worldId = worldId
	return c
}

func (c *characterAttributeBuilder) SetName(name string) *characterAttributeBuilder {
	c.name = name
	return c
}

func (c *characterAttributeBuilder) SetGender(gender byte) *characterAttributeBuilder {
	c.gender = gender
	return c
}

func (c *characterAttributeBuilder) SetSkinColor(skinColor byte) *characterAttributeBuilder {
	c.skinColor = skinColor
	return c
}

func (c *characterAttributeBuilder) SetFace(face uint32) *characterAttributeBuilder {
	c.face = face
	return c
}

func (c *characterAttributeBuilder) SetHair(hair uint32) *characterAttributeBuilder {
	c.hair = hair
	return c
}

func (c *characterAttributeBuilder) SetLevel(level byte) *characterAttributeBuilder {
	c.level = level
	return c
}

func (c *characterAttributeBuilder) SetJobId(jobId uint16) *characterAttributeBuilder {
	c.jobId = jobId
	return c
}

func (c *characterAttributeBuilder) SetStrength(strength uint16) *characterAttributeBuilder {
	c.strength = strength
	return c
}

func (c *characterAttributeBuilder) SetDexterity(dexterity uint16) *characterAttributeBuilder {
	c.dexterity = dexterity
	return c
}

func (c *characterAttributeBuilder) SetIntelligence(intelligence uint16) *characterAttributeBuilder {
	c.intelligence = intelligence
	return c
}

func (c *characterAttributeBuilder) SetLuck(luck uint16) *characterAttributeBuilder {
	c.luck = luck
	return c
}

func (c *characterAttributeBuilder) SetHp(hp uint16) *characterAttributeBuilder {
	c.hp = hp
	return c
}

func (c *characterAttributeBuilder) SetMaxHp(maxHp uint16) *characterAttributeBuilder {
	c.maxHp = maxHp
	return c
}

func (c *characterAttributeBuilder) SetMp(mp uint16) *characterAttributeBuilder {
	c.mp = mp
	return c
}

func (c *characterAttributeBuilder) SetMaxMp(maxMp uint16) *characterAttributeBuilder {
	c.maxMp = maxMp
	return c
}

func (c *characterAttributeBuilder) SetAp(ap uint16) *characterAttributeBuilder {
	c.ap = ap
	return c
}

func (c *characterAttributeBuilder) SetSp(sp string) *characterAttributeBuilder {
	c.sp = sp
	return c
}

func (c *characterAttributeBuilder) SetExperience(experience uint32) *characterAttributeBuilder {
	c.experience = experience
	return c
}

func (c *characterAttributeBuilder) SetFame(fame int16) *characterAttributeBuilder {
	c.fame = fame
	return c
}

func (c *characterAttributeBuilder) SetGachaponExperience(gachaponExperience uint32) *characterAttributeBuilder {
	c.gachaponExperience = gachaponExperience
	return c
}

func (c *characterAttributeBuilder) SetMapId(mapId uint32) *characterAttributeBuilder {
	c.mapId = mapId
	return c
}

func (c *characterAttributeBuilder) SetSpawnPoint(spawnPoint byte) *characterAttributeBuilder {
	c.spawnPoint = spawnPoint
	return c
}

func (c *characterAttributeBuilder) SetGm(gm bool) *characterAttributeBuilder {
	c.gm = gm
	return c
}

func (c *characterAttributeBuilder) SetGmJob(gmJob bool) *characterAttributeBuilder {
	c.gmJob = gmJob
	return c
}

func (c *characterAttributeBuilder) SetRank(rank int) *characterAttributeBuilder {
	c.rank = rank
	return c
}

func (c *characterAttributeBuilder) SetRankMove(rankMove int) *characterAttributeBuilder {
	c.rankMove = rankMove
	return c
}

func (c *characterAttributeBuilder) SetJobRank(jobRank int) *characterAttributeBuilder {
	c.jobRank = jobRank
	return c
}

func (c *characterAttributeBuilder) SetJobRankMove(jobRankMove int) *characterAttributeBuilder {
	c.jobRankMove = jobRankMove
	return c
}

func (c *characterAttributeBuilder) Build() Attributes {
	return Attributes{
		id:                 c.id,
		worldId:            c.worldId,
		name:               c.name,
		gender:             c.gender,
		skinColor:          c.skinColor,
		face:               c.face,
		hair:               c.hair,
		level:              c.level,
		jobId:              c.jobId,
		strength:           c.strength,
		dexterity:          c.dexterity,
		intelligence:       c.intelligence,
		luck:               c.luck,
		hp:                 c.hp,
		maxHp:              c.maxHp,
		mp:                 c.mp,
		maxMp:              c.maxMp,
		ap:                 c.ap,
		sp:                 c.sp,
		experience:         c.experience,
		fame:               c.fame,
		gachaponExperience: c.gachaponExperience,
		mapId:              c.mapId,
		spawnPoint:         c.spawnPoint,
		gm:                 c.gm,
		gmJob:              c.gmJob,
		rank:               c.rank,
		rankMove:           c.rankMove,
		jobRank:            c.jobRank,
		jobRankMove:        c.jobRankMove,
	}
}
