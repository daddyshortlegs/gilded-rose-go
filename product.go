package main

type Product interface {
	Update(*Item)
}

func lookupProduct(item *Item) interface{ Product } {
	products := createProducts()
	if i, ok := products[item.name]; ok {
		return i
	}
	return Normal{}
}

func createProducts() map[string]interface{ Product } {
	products := make(map[string]interface{ Product })
	products["Aged Brie"] = AgedBrie{}
	products["Backstage passes to a TAFKAL80ETC concert"] = Backstage{}
	products["Sulfuras, Hand of Ragnaros"] = Sulfuras{}
	return products
}
