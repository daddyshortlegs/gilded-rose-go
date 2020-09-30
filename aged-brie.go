package main

type AgedBrie struct {
	item Item
}

func (agedBrie AgedBrie) Update(item *Item) {
	if item.quality < 50 {
		item.quality++
	}
	item.sellIn--
	if item.sellIn < 0 && item.quality < 50 {
		item.quality++
	}
}
