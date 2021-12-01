package src

type BaseItem struct {
	Name            string
	SellIn, Quality int
}

func (i *BaseItem) Update() {
	i.SellIn--

	if i.SellIn > 0 {
		i.Quality -= 1
	} else {
		i.Quality -= 2
	}
}
