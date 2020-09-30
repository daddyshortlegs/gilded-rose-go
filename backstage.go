package main

type Backstage struct {
	item Item
}

func (backStage Backstage)Update(item *Item) {
	if item.quality < 50 {
		item.quality++
		if item.sellIn < 11 {
			item.quality++
		}
		if item.sellIn < 6 {
			item.quality++
		}
	}
	item.sellIn--
	if item.sellIn < 0 {
		item.quality = 0
	}
}
