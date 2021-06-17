package inventory

func GetEquippedItemsForCharacter(characterId uint32) ([]EquippedItem, error) {
	r, err := requestEquippedItemsForCharacter(characterId)
	if err != nil {
		return nil, err
	}

	eis := make([]EquippedItem, 0)
	for _, e := range r.GetIncludedEquippedItems() {
		ea := r.GetEquipmentStatistics(e.Attributes.EquipmentId)
		if ea != nil {
			ei := NewEquippedItem(ea.ItemId, e.Attributes.Slot)
			eis = append(eis, ei)
		}
	}

	return eis, nil
}