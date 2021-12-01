package src

type BaseItem struct {
	*Item
}

func NewBaseItem(item *Item) *BaseItem {
	return &BaseItem{item}
}

func (i *BaseItem) Update() {
	i.SellIn--

	if i.SellIn > 0 {
		i.Quality -= 1
	} else {
		i.Quality -= 2
	}
}
