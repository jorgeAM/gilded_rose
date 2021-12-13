package src

type Item struct {
	Name            string
	SellIn, Quality int
}

type updatableItemCreatorMap map[string]updatableItemCreator

func mapItems(items []*Item) {
	var mapper updatableItemCreatorMap = map[string]updatableItemCreator{
		"Aged Brie": func(item *Item) Updater {
			return NewAgeBrieItem(item)
		},
		"+5 Dexterity Vest": func(item *Item) Updater {
			return NewBaseItem(item)
		},
		"Elixir of the Mongoose": func(item *Item) Updater {
			return NewBaseItem(item)
		},
		"Sulfuras, Hand of Ragnaros": func(item *Item) Updater {
			return NewSulfurasItem(item)
		},
		"Backstage passes to a TAFKAL80ETC concert": func(item *Item) Updater {
			return NewBackstagePassesItem(item)
		},
		"Conjured Mana Cake": func(item *Item) Updater {
			return NewConjuredItem(item)
		},
	}

	for _, item := range items {
		updatableItemFunc := mapper[item.Name]

		updatableItem := updatableItemFunc(item)

		updatableItem.Update()
	}

}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].Quality > 0 {
				if items[i].Name != "Sulfuras, Hand of Ragnaros" {
					items[i].Quality = items[i].Quality - 1
				}
			}
		} else {
			if items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
				if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].SellIn < 11 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
					if items[i].SellIn < 6 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
				}
			}
		}

		if items[i].Name != "Sulfuras, Hand of Ragnaros" {
			items[i].SellIn = items[i].SellIn - 1
		}

		if items[i].SellIn < 0 {
			if items[i].Name != "Aged Brie" {
				if items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].Quality > 0 {
						if items[i].Name != "Sulfuras, Hand of Ragnaros" {
							items[i].Quality = items[i].Quality - 1
						}
					}
				} else {
					items[i].Quality = items[i].Quality - items[i].Quality
				}
			} else {
				if items[i].Quality < 50 {
					items[i].Quality = items[i].Quality + 1
				}
			}
		}
	}

	// mapItems(items)

}
