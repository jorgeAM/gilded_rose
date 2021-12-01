package src

type AgeBrieItem struct {
	BaseItem
}

func (i *AgeBrieItem) Update() {
	i.SellIn--

	if i.Quality < 50 {
		i.Quality++
	}
}
