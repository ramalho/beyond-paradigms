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
	return fmt.Sprintf("%d %s @ $%0.2f ea. = $%0.2f",
		li.quantity, li.product, float64(li.price)/100,
		float64(li.total())/100)
}

type FidelityPromo struct {}

// 5% discount for customers with at least 1000 fidelity points
func (p FidelityPromo) Discount(o Order) int {
	if o.customer.fidelity >= 1000 {
		return o.total() / 20  // %5
	}
	return 0
}

type BulkItemPromo struct {}

// 10% discount for each line item with at least 20 units
func (p BulkItemPromo) Discount(o Order) int {
	discount := 0
	for _, item := range o.cart {
		if item.quantity >= 20 {
			discount += item.total() / 10
		}
	}
	return discount
}

type LargeOrderPromo struct {}

// 7% discount for orders with 10 or more items
func (p LargeOrderPromo) Discount(o Order) int {
	if len(o.cart) >= 10 {
		return o.total() * 7 / 100
	}
	return 0
}
