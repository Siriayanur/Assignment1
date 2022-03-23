package controller

import (
	"fmt"
	"testing"

	"github.com/Siriayanur/Nuclei_Assignments/model"
)

func TestCalculateFinalPriceHappyFlow(t *testing.T) {

	fmt.Println("\n\n| Test | CalculateFinalPrice | Valid cases")

	itemNames := []string{"", "Wood_", "Char12oal", "Pens", "s"}
	itemPrices := []string{"120.12", "", "140000", ".", "0"}
	itemQuantities := []string{"1", "", "120", "0", "90"}
	itemTypes := []string{"1", "1", "2", "3", "3"}

	for i, name := range itemNames {
		item := model.GetItem(name, itemPrices[i], itemQuantities[i], itemTypes[i])

		if item.Error == "" {
			CalculateFinalPrice(&item)
			fmt.Println(item)
		} else {
			fmt.Println(item.Error)
		}
	}
}

//go test -v --cover
