package main

// Order is the context where a promotion discount strategy is used
type Order struct {
	customer  Customer
	cart      []LineItem
	promotion func (Order) int
}

func (o Order) total() int {
	total := 0
	for _, item := range o.cart {
		total += item.total()
	}
	return total
}

func (o Order) due() int {
	return o.total() - o.promotion(o)
}
