package main

import (
	"fmt"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// // **** Struct pointers are automatically dereferenced in Go-Lang ****

// format the bill
func (b *bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "total:", total+b.tip)

	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
	// Go will automatically convert the above line to "(*b).tip = tip"
}

// add item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
	// Go will automatically convert the above line to "(*b).items[name] = price"
}
