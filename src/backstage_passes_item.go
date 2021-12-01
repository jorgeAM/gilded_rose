package src

type BackstagePassesItem struct {
	BaseItem
}

func (i *BackstagePassesItem) Update() {
	i.SellIn--

	remainingDays := i.SellIn

	switch {
	case remainingDays > 10:
		i.Quality++

	case remainingDays <= 10 && remainingDays > 6:
		i.Quality += 2

	case remainingDays <= 5 && remainingDays > 0:
		i.Quality += 3

	case remainingDays <= 0:
		i.Quality = 0
	}
}
