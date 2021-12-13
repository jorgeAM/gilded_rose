package src

type Updater interface {
	Update()
}

type updatableItemCreator func(item *Item) Updater
