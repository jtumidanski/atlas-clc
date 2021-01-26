package models

type EquippedItem struct {
	itemId uint32
	slot   int16
}

func (i EquippedItem) NotInWeaponSlot() bool {
	if i.slot != -111 {
		return true
	}
	return false
}

func (i EquippedItem) InvertSlot() *EquippedItem {
	return NewEquippedItem(i.itemId, i.slot*-1)
}

func (i EquippedItem) Slot() int16 {
	return i.slot
}

func (i EquippedItem) ItemId() uint32 {
	return i.itemId
}

func (i EquippedItem) InWeaponSlot() bool {
	if i.slot == -111 {
		return true
	}
	return false
}

func NewEquippedItem(itemId uint32, slot int16) *EquippedItem {
	return &EquippedItem{itemId, slot}
}
