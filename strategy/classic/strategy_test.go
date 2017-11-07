package main

import (
	"fmt"
	"testing"
)

var joe = Customer{"John Doe", 0}
var bananas = LineItem{"banana", 4, 50}

func TestCustomer(t *testing.T) {
	name := "John Doe"
	if joe.name != name {
		t.Errorf("want: %q; got: %q", name, joe.name)
	}
}

func TestLineItem(t *testing.T) {
	want := "4 banana @ $0.50 = $2.00"
	got := fmt.Sprint(bananas)
	if want != got {
		t.Errorf("\nwant: %q;\n got: %q", want, got)
	}
}
