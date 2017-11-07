package main

import (
	"fmt"
)

// Customer has name and fidelity points
type Customer struct {
	name     string
	fidelity int
}

// LineItem represents one item in an Order
type LineItem struct {
	product  string
	quantity int
	price    int // cents
}

func (li LineItem) total() int {
	return li.quantity * li.price
}

func (li LineItem) String() string {
	return fmt.Sprintf("%d %s @ $%0.2f = $%0.2f",
		li.quantity, li.product, float64(li.price)/100,
		float64(li.total())/100)
}
