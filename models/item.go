package models

type Item struct {
	count              uint
	ChargeCount        uint
	Fluid              rune
	ActionId           uint
	UniqueId           uint
	SpecialDescription string
	Text               string
	Items              Items
}

func (i *Item) IsStackable() bool {
	return i.Items.Stackable
}
