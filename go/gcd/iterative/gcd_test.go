package intmath

import (
	"fmt"
	"testing"
)

func TestGcd(t *testing.T) {
	want := 3
	got := Gcd(6, 9)
	if got != want {
		t.Errorf("Gcd(6, 9) = %d; want %d", got, want)
	}
}

func TestGcd_table(t *testing.T) {
	testCases := []struct {
		a    int
		b    int
		want int
	}{
		{6, 9, 3},
		{9, 6, 3},
		{2, 7, 1},
		{42, 28, 14},
		{42, 41, 1},
		{42, 42, 42},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%x:%x", tc.a, tc.b), func(t *testing.T) {
			got := Gcd(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("Gcd(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
