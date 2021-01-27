package processors

import (
	"atlas-clc/models"
	"atlas-clc/rest/attributes"
	"atlas-clc/rest/requests"
	"errors"
	"log"
	"regexp"
	"strconv"
)

func GetCharacterAttributesByName(l *log.Logger, name string) (*models.CharacterAttributes, error) {
	ca, err := requests.GetCharacterAttributesByName(l, name)
	if err != nil {
		return nil, err
	}
	if len(ca.DataList()) <= 0 {
		return nil, errors.New("unable to find character by name")
	}

	return makeCharacterAttributes(ca.Data()), nil
}

func makeCharacterAttributes(ca *attributes.CharacterAttributesData) *models.CharacterAttributes {
	cid, err := strconv.ParseUint(ca.Id, 10, 32)
	if err != nil {
		return nil
	}
	att := ca.Attributes
	return models.NewCharacterAttributes(uint32(cid), att.WorldId, att.Name, att.Gender, att.SkinColor, att.Face, att.Hair, att.Level, att.JobId, att.Strength, att.Dexterity, att.Intelligence, att.Luck, att.Hp, att.MaxHp, att.Mp, att.MaxMp, att.Ap, att.Sp, att.Experience, att.Fame, att.GachaponExperience, att.MapId, att.SpawnPoint)
}

func IsValidName(l *log.Logger, name string) (bool, error) {
	m, err := regexp.MatchString("[a-zA-Z0-9]{3,12}", name)
	if err != nil {
		l.Println("[ERROR] error processing regex for character name matching")
		return false, err
	}
	if !m {
		return false, nil
	}

	_, err = GetCharacterAttributesByName(l, name)
	if err == nil {
		return false, nil
	}

	if err.Error() != "unable to find character by name" {
		return false, nil
	}

	bn, err := IsBlockedName(l, name)
	if bn {
		return false, err
	}

	return true, nil
}

func GetCharactersForWorld(l *log.Logger, accountId int, worldId byte) ([]models.Character, error) {
	cs, err := requests.GetCharacterAttributesForAccountByWorld(l, accountId, worldId)
	if err != nil {
		l.Println("[ERROR] error retrieving characters for account by world")
		return nil, err
	}

	var characters = make([]models.Character, 0)
	for _, x := range cs.DataList() {
		c, err := getCharacterForAttributes(l, &x)
		if err != nil {
			l.Println("[ERROR] error retrieving character detail")
			return nil, err
		}
		characters = append(characters, *c)
	}
	return characters, nil
}

func GetCharacterById(l *log.Logger, characterId uint32) (*models.Character, error) {
	cs, err := requests.GetCharacterAttributesById(l, characterId)
	if err != nil {
		l.Println("[ERROR] error retrieving character by id")
		return nil, err
	}

	c, err := getCharacterForAttributes(l, cs.Data())
	if err != nil {
		return nil, err
	}
	return c, nil
}

func getCharacterForAttributes(l *log.Logger, data *attributes.CharacterAttributesData) (*models.Character, error) {
	ca := makeCharacterAttributes(data)
	if ca == nil {
		return nil, errors.New("unable to make character attributes")
	}

	eq, err := getEquippedItemsForCharacter(l, ca.Id())
	if err != nil {
		return nil, err
	}

	ps, err := getPetsForCharacter(l, ca.Id())
	if err != nil {
		return nil, err
	}

	return models.NewCharacter(*ca, eq, ps), nil
}

func getPetsForCharacter(l *log.Logger, characterId uint32) ([]models.Pet, error) {
	return make([]models.Pet, 0), nil
}

func getEquippedItemsForCharacter(l *log.Logger, characterId uint32) ([]models.EquippedItem, error) {
	r, err := requests.GetEquippedItemsForCharacter(l, characterId)
	if err != nil {
		return nil, err
	}

	ei := make([]models.EquippedItem, 0)
	for _, e := range r.GetIncludedEquippedItems() {
		ea := r.GetEquipmentStatistics(e.Attributes.EquipmentId)
		if ea != nil {
			ei = append(ei, *models.NewEquippedItem(ea.ItemId, e.Attributes.Slot))
		}
	}

	return ei, nil
}

func SeedCharacter(l *log.Logger, accountId int, worldId byte, name string, job uint32, face uint32, hair uint32, color uint32, skinColor uint32, gender byte, top uint32, bottom uint32, shoes uint32, weapon uint32) (*models.CharacterAttributes, error) {
	ca, err := requests.SeedCharacter(l, accountId, worldId, name, job, face, hair, color, skinColor, gender, top, bottom, shoes, weapon)
	if err != nil {
		return nil, err
	}
	return makeCharacterAttributes(ca), nil
}
