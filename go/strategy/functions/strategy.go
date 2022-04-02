package main

// Order is the context where a promotion discount strategy is used
type Order struct {
	Customer  Customer
	Cart      []LineItem
	Promotion func(Order) int
}

func (o Order) Total() int {
	total := 0
	for _, item := range o.Cart {
		total += item.Total()
	}
	return total
}

func (o Order) Due() int {
	return o.Total() - o.Promotion(o)
}
