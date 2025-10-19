package main

import "fmt"

// Inventory tracks available items in the restaurant
type Inventory struct {
	stock map[string]int
}

func newInventory() *Inventory {
	return &Inventory{
		stock: map[string]int{
			"Pizza": 10,
			"Pasta": 15,
			"Salad": 20,
		},
	}
}

func (i *Inventory) checkStock(item string, quantity int) error {
	available, exists := i.stock[item]
	if !exists {
		return fmt.Errorf("item '%s' not on menu", item)
	}
	if available < quantity {
		return fmt.Errorf("only %d x %s available", available, item)
	}
	fmt.Printf("Stock confirmed: %d x %s available\n", available, item)
	return nil
}

func (i *Inventory) reduceStock(item string, quantity int) {
	i.stock[item] -= quantity
	fmt.Printf("Stock updated: %d x %s remaining\n", i.stock[item], item)
}