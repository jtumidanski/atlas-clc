package writers

import (
	"atlas-clc/models"
	"atlas-clc/packets/outputs"
	"atlas-clc/packets/outputs/constants"
)

func WriteCharacterList(characters []models.Character, worldId byte, status int, cannotBypassPic bool, pic string, availableCharacterSlots int16, characterSlots int16) []byte {
	w := outputs.NewWriter()
	w.WriteShort(constants.CharacterList)
	w.WriteByte(byte(status))
	w.WriteByte(byte(len(characters)))
	for _, x := range characters {
		WriteCharacter(w, x, false)
	}
	w.WriteByte(2)
	w.WriteInt(uint32(characterSlots))
	return w.Bytes()
}

func WriteCharacter(w *outputs.Writer, character models.Character, viewAll bool) {
	WriteCharacterStatistics(w, character)
	WriteCharacterLook(w, character, false)
	if !viewAll {
		w.WriteByte(0)
	}
	if character.Attributes().Gm() || character.Attributes().GmJob() {
		w.WriteByte(0)
		return
	}
	w.WriteByte(1) // world rank enabled (next 4 ints are not sent if disabled) Short??
	w.WriteInt(uint32(character.Attributes().Rank()))
	w.WriteInt(uint32(character.Attributes().RankMove()))
	w.WriteInt(uint32(character.Attributes().JobRank()))
	w.WriteInt(uint32(character.Attributes().JobRankMove()))
}

func WriteCharacterLook(w *outputs.Writer, character models.Character, mega bool) {
	w.WriteByte(character.Attributes().Gender())
	w.WriteByte(character.Attributes().SkinColor())
	w.WriteInt(character.Attributes().Face())
	w.WriteBool(!mega)
	w.WriteInt(character.Attributes().Hair())
	WriteCharacterEquipment(w, character)
}

func WriteCharacterEquipment(w *outputs.Writer, character models.Character) {

	var equips = getEquipedItemSlotMap(character.Equipment())
	var maskedEquips = make(map[int16]uint32)
	writeEquips(w, equips, maskedEquips)

	var weapon *models.EquippedItem
	for _, x := range character.Equipment() {
		if x.InWeaponSlot() {
			weapon = &x
			break
		}
	}
	if weapon != nil {
		w.WriteInt(weapon.ItemId())
	} else {
		w.WriteInt(0)
	}

	writeForEachPet(w, character.Pets(), writePetItemId, writeEmptyPetItemId)
}

func writeEquips(w *outputs.Writer, equips map[int16]uint32, maskedEquips map[int16]uint32) {
	for k, v := range equips {
		w.WriteKeyValue(byte(k), v)
	}
	w.WriteByte(0xFF)
	for k, v := range maskedEquips {
		w.WriteKeyValue(byte(k), v)
	}
	w.WriteByte(0xFF)
}

func getEquipedItemSlotMap(e []models.EquippedItem) map[int16]uint32 {
	var equips = make(map[int16]uint32, 0)
	for _, x := range e {
		if x.NotInWeaponSlot() {
			y := x.InvertSlot()
			equips[y.Slot()] = y.ItemId()
		}
	}
	return equips
}

func writePetItemId(w *outputs.Writer, p models.Pet) {
	w.WriteInt(p.ItemId())
}

func writeEmptyPetItemId(w *outputs.Writer) {
	w.WriteInt(0)
}

func writeForEachPet(w *outputs.Writer, ps []models.Pet, pe func(w *outputs.Writer, p models.Pet), pne func(w *outputs.Writer)) {
	for i := 0; i < 3; i++ {
		if ps != nil && len(ps) > i {
			pe(w, ps[i])
		} else {
			pne(w)
		}
	}
}

func writePetId(w *outputs.Writer, pet models.Pet) {
	w.WriteLong(pet.Id())
}

func writeEmptyPetId(w *outputs.Writer) {
	w.WriteLong(0)
}

func WriteCharacterStatistics(w *outputs.Writer, character models.Character) {
	w.WriteInt(character.Attributes().Id())

	name := character.Attributes().Name()
	if len(name) > 13 {
		name = name[:13]
	}
	padSize := 13 - len(name)
	w.WriteByteArray([]byte(name))
	for i := 0; i < padSize; i++ {
		w.WriteByte(0x0)
	}

	w.WriteByte(character.Attributes().Gender())
	w.WriteByte(character.Attributes().SkinColor())
	w.WriteInt(character.Attributes().Face())
	w.WriteInt(character.Attributes().Hair())
	writeForEachPet(w, character.Pets(), writePetId, writeEmptyPetId)
	w.WriteByte(character.Attributes().Level())
	w.WriteShort(character.Attributes().JobId())
	w.WriteShort(character.Attributes().Strength())
	w.WriteShort(character.Attributes().Dexterity())
	w.WriteShort(character.Attributes().Intelligence())
	w.WriteShort(character.Attributes().Luck())
	w.WriteShort(character.Attributes().Hp())
	w.WriteShort(character.Attributes().MaxHp())
	w.WriteShort(character.Attributes().Mp())
	w.WriteShort(character.Attributes().MaxMp())
	w.WriteShort(character.Attributes().Ap())

	if character.Attributes().HasSPTable() {
		WriteRemainingSkillInfo(w, character)
	} else {
		w.WriteShort(character.Attributes().RemainingSp())
	}

	w.WriteInt(character.Attributes().Experience())
	w.WriteShort(uint16(character.Attributes().Fame()))
	w.WriteInt(character.Attributes().GachaponExperience())
	w.WriteInt(character.Attributes().MapId())
	w.WriteByte(character.Attributes().SpawnPoint())
	w.WriteInt(0)
}

func WriteRemainingSkillInfo(w *outputs.Writer, character models.Character) {

}