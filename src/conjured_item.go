package src

type ConjuredItem struct {
	*Item
}

func NewConjuredItem(item *Item) *ConjuredItem {
	return &ConjuredItem{item}
}

func (i *ConjuredItem) Update() {
	i.SellIn--

	i.Quality -= 2
}
