package src

import "fmt"

type SulfurasItem struct {
	*Item
}

func NewSulfurasItem(item *Item) *SulfurasItem {
	return &SulfurasItem{item}
}

func (i *SulfurasItem) Update() {
	fmt.Println("Sulfuras item are legendary, so nothing happens")
}
