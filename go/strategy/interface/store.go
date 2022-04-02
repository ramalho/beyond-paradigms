package main

import (
	"fmt"
)

// Customer has name and fidelity points
type Customer struct {
	Name     string
	Fidelity int
}

// LineItem represents one item in an Order
type LineItem struct {
	Product  string
	Quantity int
	Price    int // cents
}

func (li LineItem) Total() int {
	return li.Quantity * li.Price
}

func (li LineItem) String() string {
	return fmt.Sprintf("%d %s @ $%0.2f ea. = $%0.2f",
		li.Quantity, li.Product, float64(li.Price)/100,
		float64(li.Total())/100)
}

type FidelityPromo struct{}

// 5% discount for customers with at least 1000 fidelity points
func (p FidelityPromo) Discount(o Order) int {
	if o.Customer.Fidelity >= 1000 {
		return o.Total() / 20 // %5
	}
	return 0
}

type BulkItemPromo struct{}

// 10% discount for each line item with at least 20 units
func (p BulkItemPromo) Discount(o Order) int {
	discount := 0
	for _, item := range o.Cart {
		if item.Quantity >= 20 {
			discount += item.Total() / 10
		}
	}
	return discount
}

type LargeOrderPromo struct{}

// 7% discount for orders with 10 or more items
func (p LargeOrderPromo) Discount(o Order) int {
	if len(o.Cart) >= 10 {
		return o.Total() * 7 / 100
	}
	return 0
}
