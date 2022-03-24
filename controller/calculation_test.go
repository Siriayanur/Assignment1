package controller

import (
	"testing"

	"github.com/Siriayanur/Nuclei_Assignments/model"
)

func TestCalculateFinalPrice(t *testing.T) {

	items := []struct {
		Name          string
		Price         string
		Quantity      string
		Types         string
		expSalesTax   float64
		expFinalPrice float64
	}{
		{Name: "Pen", Price: "50", Quantity: "3", Types: "1", expSalesTax: 6.2500, expFinalPrice: 168.7500},
		{Name: "BRICK", Price: "128", Quantity: "10", Types: "1", expSalesTax: 16.0000, expFinalPrice: 1440.00},
		{Name: "Glass_", Price: "300.50", Quantity: "1", Types: "2", expSalesTax: 44.3238, expFinalPrice: 344.8238},
		{Name: "", Price: "10.0", Quantity: "4", Types: "3", expSalesTax: 5.1000, expFinalPrice: 60.4000},
		{Name: "chocs123", Price: "-190.0", Quantity: "10", Types: "2", expSalesTax: 0.0, expFinalPrice: 0.0},
		{Name: "@##", Price: "190.0", Quantity: "10", Types: "2", expSalesTax: 0.0, expFinalPrice: 0.0},
		{Name: "hello", Price: "190.0", Quantity: "-9", Types: "3", expSalesTax: 0.0, expFinalPrice: 0.0},
		{Name: "1234", Price: "80.0", Quantity: "10", Types: "4", expSalesTax: 0.0, expFinalPrice: 0.0},
		{Name: "table", Price: "100", Quantity: "9.9", Types: "2", expSalesTax: 0.0, expFinalPrice: 0.0},
	}

	for _, item := range items {
		validItem, err := model.GetItem(item.Name, item.Price, item.Quantity, item.Types)
		if err == "" {
			CalculateFinalPrice(&validItem)
			if validItem.FinalPrice != item.expFinalPrice || validItem.SalesTax != item.expSalesTax {
				t.Errorf("Expected Final Price is %v, but got %v\n", item.expFinalPrice, validItem.FinalPrice)
			}
		} else {
			t.Errorf("Incorrect Item parameter : %s", err)
		}
	}
}
