package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Example() {
	sample := []float64{0, 1, -2, 3, -4, 5}
	for x := range PositiveIterator(sample) {
		fmt.Printf("%0.1f ", x)
	}
	// Output: 1.0 3.0 5.0
}

func Test_scan(t *testing.T) {
	testCases := []struct {
		given []float64
		want  []float64
	}{
		{[]float64{}, []float64{}},
		{[]float64{0}, []float64{}},
		{[]float64{1}, []float64{1}},
		{[]float64{1, 2, 3}, []float64{1, 2, 3}},
		{[]float64{-1, 2, 3}, []float64{2, 3}},
		{[]float64{1, -2, 3}, []float64{1, 3}},
		{[]float64{1, 2, -3}, []float64{1, 2}},
		{[]float64{-1, -2, 3}, []float64{3}},
		{[]float64{-1, -2, -3}, []float64{}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.given), func(t *testing.T) {
			got := make([]float64, 0, len(tc.given))
			for v := range PositiveIterator(tc.given) {
				got = append(got, v)
			}
			assert.Equal(t, tc.want, got)
		})
	}
}
