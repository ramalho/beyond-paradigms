package main

import (
	"fmt"
	"testing"
)

var bananas = LineItem{"banana", 4, 50}
var apples = LineItem{"apple", 10, 150}
var mellons = LineItem{"mellon", 5, 500}

func TestLineItem_String(t *testing.T) {
	want := "4 banana @ $0.50 ea. = $2.00"
	got := fmt.Sprint(bananas)
	if want != got {
		t.Errorf("\nwant: %q;\n got: %q", want, got)
	}
}

var joe = Customer{"John Doe", 0}

func TestOrderTotal(t *testing.T) {
	cart := []LineItem{bananas, apples, mellons}
	order := Order{joe, cart, FidelityPromo{}}
	want := 4200
	got := order.Total()
	if want != got {
		t.Errorf("\nwant: %v;\n total: %v", want, got)
	}
}

func TestNoFidelityOrder(t *testing.T) {
	cart := []LineItem{bananas, apples, mellons}
	order := Order{joe, cart, FidelityPromo{}}
	want := 4200
	got := order.Due()
	if want != got {
		t.Errorf("\nwant: %v;\n got: %v", want, got)
	}
}

func TestFidelityOrder(t *testing.T) {
	customer := Customer{"Ann Smith", 1000}
	cart := []LineItem{bananas, apples, mellons}
	order := Order{customer, cart, FidelityPromo{}}
	want := 3990
	got := order.Due()
	if want != got {
		t.Errorf("\nwant: %v;\n got: %v", want, got)
	}
}

func TestBulkItemOrder(t *testing.T) {
	cart := []LineItem{{"banana", 30, 50}, apples}
	order := Order{joe, cart, BulkItemPromo{}}
	want := 2850
	got := order.Due()
	if want != got {
		t.Errorf("\nwant: %v;\n got: %v", want, got)
	}
}

func TestLargeOrder(t *testing.T) {
	cart := []LineItem{}
	for i := 0; i < 10; i++ {
		cart = append(cart, LineItem{string(i), 1, 100})
	}
	order := Order{joe, cart, LargeOrderPromo{}}
	want := 930
	got := order.Due()
	if want != got {
		t.Errorf("\nwant: %v;\n got: %v", want, got)
	}
}
