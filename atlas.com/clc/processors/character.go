package processors

import (
   "atlas-clc/domain"
   "atlas-clc/rest/attributes"
   "atlas-clc/rest/requests"
   "errors"
   "regexp"
   "strconv"
)

func GetCharacterAttributesByName(name string) (*domain.CharacterAttributes, error) {
   ca, err := requests.GetCharacterAttributesByName(name)
   if err != nil {
      return nil, err
   }
   if len(ca.DataList()) <= 0 {
      return nil, errors.New("unable to find character by name")
   }

   return makeCharacterAttributes(ca.Data()), nil
}

func makeCharacterAttributes(ca *attributes.CharacterAttributesData) *domain.CharacterAttributes {
   cid, err := strconv.ParseUint(ca.Id, 10, 32)
   if err != nil {
      return nil
   }
   att := ca.Attributes
   r := domain.NewCharacterAttributeBuilder().
      SetId(uint32(cid)).
      SetWorldId(att.WorldId).
      SetName(att.Name).
      SetGender(att.Gender).
      SetSkinColor(att.SkinColor).
      SetFace(att.Face).
      SetHair(att.Hair).
      SetLevel(att.Level).
      SetJobId(att.JobId).
      SetStrength(att.Strength).
      SetDexterity(att.Dexterity).
      SetIntelligence(att.Intelligence).
      SetLuck(att.Luck).
      SetHp(att.Hp).
      SetMaxHp(att.MaxHp).
      SetMp(att.Mp).
      SetMaxMp(att.MaxMp).
      SetAp(att.Ap).
      SetSp(att.Sp).
      SetExperience(att.Experience).
      SetFame(att.Fame).
      SetGachaponExperience(att.GachaponExperience).
      SetMapId(att.MapId).
      SetSpawnPoint(att.SpawnPoint).
      Build()
   return &r
}

func IsValidName(name string) (bool, error) {
   m, err := regexp.MatchString("[a-zA-Z0-9]{3,12}", name)
   if err != nil {
      return false, err
   }
   if !m {
      return false, nil
   }

   _, err = GetCharacterAttributesByName(name)
   if err == nil {
      return false, nil
   }

   if err.Error() != "unable to find character by name" {
      return false, nil
   }

   bn, err := IsBlockedName(name)
   if bn {
      return false, err
   }

   return true, nil
}

func GetCharactersForWorld(accountId uint32, worldId byte) ([]domain.Character, error) {
   cs, err := requests.GetCharacterAttributesForAccountByWorld(accountId, worldId)
   if err != nil {
      return nil, err
   }

   var characters = make([]domain.Character, 0)
   for _, x := range cs.DataList() {
      c, err := getCharacterForAttributes(&x)
      if err != nil {
         return nil, err
      }
      characters = append(characters, *c)
   }
   return characters, nil
}

func GetCharacterById(characterId uint32) (*domain.Character, error) {
   cs, err := requests.GetCharacterAttributesById(characterId)
   if err != nil {
      return nil, err
   }

   c, err := getCharacterForAttributes(cs.Data())
   if err != nil {
      return nil, err
   }
   return c, nil
}

func getCharacterForAttributes(data *attributes.CharacterAttributesData) (*domain.Character, error) {
   ca := makeCharacterAttributes(data)
   if ca == nil {
      return nil, errors.New("unable to make character attributes")
   }

   eq, err := getEquippedItemsForCharacter(ca.Id())
   if err != nil {
      return nil, err
   }

   ps, err := getPetsForCharacter()
   if err != nil {
      return nil, err
   }

   c := domain.NewCharacter(*ca, eq, ps)
   return &c, nil
}

func getPetsForCharacter() ([]domain.Pet, error) {
   return make([]domain.Pet, 0), nil
}

func getEquippedItemsForCharacter(characterId uint32) ([]domain.EquippedItem, error) {
   r, err := requests.GetEquippedItemsForCharacter(characterId)
   if err != nil {
      return nil, err
   }

   eis := make([]domain.EquippedItem, 0)
   for _, e := range r.GetIncludedEquippedItems() {
      ea := r.GetEquipmentStatistics(e.Attributes.EquipmentId)
      if ea != nil {
         ei := domain.NewEquippedItem(ea.ItemId, e.Attributes.Slot)
         eis = append(eis, ei)
      }
   }

   return eis, nil
}

func SeedCharacter(accountId uint32, worldId byte, name string, job uint32, face uint32, hair uint32, color uint32, skinColor uint32, gender byte, top uint32, bottom uint32, shoes uint32, weapon uint32) (*domain.CharacterAttributes, error) {
   ca, err := requests.SeedCharacter(accountId, worldId, name, job, face, hair, color, skinColor, gender, top, bottom, shoes, weapon)
   if err != nil {
      return nil, err
   }
   return makeCharacterAttributes(ca), nil
}
