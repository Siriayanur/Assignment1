package model

import (
	"testing"
)

func TestItemNameValidCases(t *testing.T) {
	itemNames := []string{
		"HeLlO",
		"PEN",
		"1234",
		"_HAT_",
		"",
	}

	for _, name := range itemNames {
		_, nameError := validateItemName(name)
		if nameError != nil {
			t.Errorf("Expected alphanumeric characters and underscore only, got : %s", name)
		}

	}

}
func TestItemNameInvalidCases(t *testing.T) {

	itemNames := []string{
		"hi@12",
		"$$@",
		"--",
		"he/<>1"}

	for _, name := range itemNames {
		_, nameError := validateItemName(name)
		if nameError == nil {
			t.Errorf("Expected alphanumeric characters and underscore only, got : %s", name)
		}

	}

}

func TestItemPriceValidCases(t *testing.T) {
	itemPrices := []string{
		"120.00",
		"120",
		"34.333",
	}

	for _, price := range itemPrices {
		_, priceError := validateItemPrice(price)
		if priceError != nil {
			t.Errorf("Expected positive float number only, got : %s", price)
		}
	}

}
func TestItemPriceInvalidCases(t *testing.T) {
	itemPrices := []string{
		".",
		"uhmm",
		"120.abc",
		"-34.333",
		"-90",
		"@#"}

	for _, price := range itemPrices {
		_, priceError := validateItemPrice(price)
		if priceError == nil {
			t.Errorf("Expected positive float number only, got : %s", price)
		}
	}

}
func TestItemQuantityValidCases(t *testing.T) {

	itemQuantity := []string{"0", "100", "13"}

	for _, quantity := range itemQuantity {
		_, quantityError := validateItemQuantity(quantity)
		if quantityError != nil {
			t.Errorf("Expected positive integer number only, got : %s", quantity)
		}
	}
}
func TestItemQuantityInvalidCases(t *testing.T) {

	itemQuantity := []string{"we", "-100", "13.45", "-56.5"}

	for _, quantity := range itemQuantity {
		_, quantityError := validateItemQuantity(quantity)
		if quantityError == nil {
			t.Errorf("Expected positive integer number only, got : %s", quantity)
		}
	}
}
func TestItemTypeValidCases(t *testing.T) {
	itemType := []string{"3", "1", "2"}

	for _, types := range itemType {
		_, typesError := validateItemType(types)
		if typesError != nil {
			t.Errorf("Expected 1, 2 or 3 only, got : %s", types)
		}
	}
}
func TestItemTypeInvalidCases(t *testing.T) {
	itemType := []string{"-3", "4", "hi", "$#", "7.9", "0"}

	for _, types := range itemType {
		_, typesError := validateItemType(types)
		if typesError == nil {
			t.Errorf("Expected 1, 2 or 3 only, got : %s", types)
		}
	}
}

func TestGetItemValidCases(t *testing.T) {
	items := []struct {
		Name     string
		Price    string
		Quantity string
		Types    string
	}{
		{Name: "Pen", Price: "50", Quantity: "3", Types: "1"},
		{Name: "BRICK", Price: "128", Quantity: "", Types: "1"},
		{Name: "Glass_", Price: "", Quantity: "1", Types: "2"},
		{Name: "", Price: "10.0", Quantity: "4", Types: "3"},
	}
	for _, item := range items {
		_, err := GetItem(item.Name, item.Price, item.Quantity, item.Types)
		if err != "" {
			t.Errorf("Incorrect Item parameters : %s", err)
		}
	}

}
func TestGetItemInvalidCases(t *testing.T) {
	items := []struct {
		Name     string
		Price    string
		Quantity string
		Types    string
	}{
		{Name: "chocs123", Price: "-190.0", Quantity: "10", Types: "2"},
		{Name: "@##", Price: "190.0", Quantity: "10", Types: "2"},
		{Name: "hello", Price: "190.0", Quantity: "-9", Types: "3"},
		{Name: "1234", Price: "80.0", Quantity: "10", Types: "4"},
		{Name: "table", Price: "100", Quantity: "9.9", Types: "2"},
	}
	for _, item := range items {
		_, err := GetItem(item.Name, item.Price, item.Quantity, item.Types)
		if err == "" {
			t.Errorf("Incorrect Item parameters : %s", err)
		}
	}

}
