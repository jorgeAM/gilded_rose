package src

type AgeBrieItem struct {
	*Item
}

func NewAgeBrieItem(item *Item) *AgeBrieItem {
	return &AgeBrieItem{item}
}

func (i *AgeBrieItem) Update() {
	i.SellIn--

	if i.Quality < 50 {
		i.Quality++
	}
}
