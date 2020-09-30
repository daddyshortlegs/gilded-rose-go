package main

import "testing"

func Test_Foo(t *testing.T) {
	var items = []*Item{
		&Item{"foo", 0, 0},
		&Item{"foo", 10, 6},
		&Item{"foo", 0, 6},

		&Item{"Aged Brie", 0, 6},
		&Item{"Aged Brie", 10, 10},
		&Item{"Aged Brie", 0, 50},

		&Item{"Sulfuras, Hand of Ragnaros", 10, 10},

		&Item{"Backstage passes to a TAFKAL80ETC concert", 10, 10},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 12, 10},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 5, 10},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 0, 10},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 12, 50},
	}

	UpdateQuality(items)

	verifyItem(t, items[0], "foo", -1, 0)
	verifyItem(t, items[1], "foo", 9, 5)
	verifyItem(t, items[2], "foo", -1, 4)
	verifyItem(t, items[3], "Aged Brie", -1, 8)
	verifyItem(t, items[4], "Aged Brie", 9, 11)
	verifyItem(t, items[5], "Aged Brie", -1, 50)
	verifyItem(t, items[6], "Sulfuras, Hand of Ragnaros", 10, 10)
	verifyItem(t, items[7], "Backstage passes to a TAFKAL80ETC concert", 9, 12)
	verifyItem(t, items[8], "Backstage passes to a TAFKAL80ETC concert", 11, 11)
	verifyItem(t, items[9], "Backstage passes to a TAFKAL80ETC concert", 4, 13)
	verifyItem(t, items[10], "Backstage passes to a TAFKAL80ETC concert", -1, 0)
	verifyItem(t, items[11], "Backstage passes to a TAFKAL80ETC concert", 11, 50)
}


func verifyItem(t *testing.T, item *Item, name string, sellIn int, quality int) {
	assertEqual(t, name, item.name)
	assertEqual(t, quality, item.quality)
	assertEqual(t, sellIn, item.sellIn)
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Errorf("Name: Expected %s but got %s ", expected, actual)
	}
}
