package main

import (
	"fmt"
	"testing"
)

func TestTaxCalculation(t *testing.T) {
	itemTypes := []int{0, 1, 2}
	itemPrices := []float64{40.0, 30.0, 100.5}
	fmt.Println(itemTypes)
	fmt.Println(itemPrices)
}
