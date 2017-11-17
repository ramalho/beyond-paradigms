package main

// Order is the context where a Discounter strategy is used
type Order struct {
	customer  Customer
	cart      []LineItem
	promotion Discounter
}

type Discounter interface {
	Discount(Order) int
}

func (o Order) total() int {
	total := 0
	for _, item := range o.cart {
		total += item.total()
	}
	return total
}

func (o Order) due() int {
	return o.total() - o.promotion.Discount(o)
}
