package src

import "fmt"

type SulfurasItem struct {
	BaseItem
}

func (i *SulfurasItem) Update() {
	fmt.Println("Sulfuras item are legendary, so nothing happens")
}
