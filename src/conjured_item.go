package src

type ConjuredItem struct {
	BaseItem
}

func (i *ConjuredItem) Update() {
	i.SellIn--

	i.Quality -= 2
}
