package inventory

import "github.com/sirupsen/logrus"

func GetEquippedItemsForCharacter(l logrus.FieldLogger) func(characterId uint32) ([]EquippedItem, error) {
	return func(characterId uint32) ([]EquippedItem, error) {
		r, err := requestEquippedItemsForCharacter(l)(characterId)
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
}
