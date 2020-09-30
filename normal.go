package main

type Normal struct {
}

func (n Normal) Update(item *Item) {
	if item.quality > 0 {
		item.quality--
	}

	item.sellIn--
	if item.quality > 0 && item.sellIn < 0 {
		item.quality--
	}
}
