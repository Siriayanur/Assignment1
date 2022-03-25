package controller

import (
	"testing"

	"github.com/Siriayanur/Nuclei_Assignments/model"
)

func TestCalculateFinalPriceValidCases(t *testing.T) {
	items := []struct {
		Name          string
		Price         string
		Quantity      string
		Types         string
		expSalesTax   float64
		expFinalPrice float64
	}{
		{Name: "Pen", Price: "50", Quantity: "3", Types: "1", expSalesTax: 6.25, expFinalPrice: 168.75},
		{Name: "BRICK", Price: "10", Quantity: "5", Types: "3", expSalesTax: 1.47, expFinalPrice: 57.35},
		{Name: "Glass_", Price: "110.50", Quantity: "1", Types: "2", expSalesTax: 6.10, expFinalPrice: 116.60},
		{Name: "", Price: "1000.0", Quantity: "4", Types: "3", expSalesTax: 15.00, expFinalPrice: 4060.00},
		{Name: "Chocs123", Price: "90.0", Quantity: "1", Types: "3", expSalesTax: 13.27, expFinalPrice: 103.27},
		{Name: "1234", Price: "110.50", Quantity: "1", Types: "3", expSalesTax: 16.29, expFinalPrice: 126.79},

		{Name: "ex1", Price: "", Quantity: "1", Types: "1", expSalesTax: 0.00, expFinalPrice: 0.00},
		{Name: "ex2", Price: "20", Quantity: "", Types: "2", expSalesTax: 5.20, expFinalPrice: 25.20},
	}
	for _, item := range items {
		validItem, _ := model.GetItem(item.Name, item.Price, item.Quantity, item.Types)
		// require.Equal(t,err, "") --> gives error
		CalculateFinalPrice(&validItem)
		if validItem.FinalPrice != item.expFinalPrice || validItem.SalesTax != item.expSalesTax {
			t.Errorf("Expected Final Price is %v, but got %v\n", item.expFinalPrice, validItem.FinalPrice)
		}
	}
}

func TestCalculateFinalPriceInvalidCases(t *testing.T) {
	items := []struct {
		Name          string
		Price         string
		Quantity      string
		Types         string
		expSalesTax   float64
		expFinalPrice float64
	}{
		{Name: "chocs123", Price: "-190.0", Quantity: "10", Types: "2", expSalesTax: 0.0, expFinalPrice: 0.0},
		{Name: "@##", Price: "190.0", Quantity: "10", Types: "2", expSalesTax: 0.0, expFinalPrice: 0.0},
		{Name: "hello", Price: "190.0", Quantity: "-9", Types: "3", expSalesTax: 0.0, expFinalPrice: 0.0},
		{Name: "1234", Price: "80.0", Quantity: "10", Types: " ", expSalesTax: 0.0, expFinalPrice: 0.0},
		{Name: "table", Price: "100", Quantity: "9.9", Types: "4", expSalesTax: 0.0, expFinalPrice: 0.0},
	}
	for _, item := range items {
		_, err := model.GetItem(item.Name, item.Price, item.Quantity, item.Types)
		if err == "" {
			t.Errorf("Incorrect Item parameter : %s", err)
		}
	}
}
