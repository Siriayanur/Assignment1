package model

import (
	"log"
	"testing"
)

func TestItemName(t *testing.T) {
	log.Println("Item Name Check")

	itemNames := []string{
		"HeLlO",
		"PEN",
		"1234",
		"_HAT_",
		"",
		"hi@12",
		"$$@",
		"--",
		"he/<>1"}

	for _, name := range itemNames {
		_, nameError := validateItemName(name)
		if nameError != nil {
			t.Errorf("Expected alphanumeric characters and underscore only, got : %s", name)
		}

	}

}

func TestItemPrice(t *testing.T) {
	log.Println("Item Price Check")
	itemPrices := []string{
		"120.00",
		"120",
		"34.333",
		".",
		"uhmm",
		"120.abc",
		"-34.333",
		"-90",
		"@#"}

	for _, price := range itemPrices {
		_, priceError := validateItemPrice(price)
		if priceError != nil {
			t.Errorf("Expected positive float number only, got : %s", price)
		}
	}

}

func TestItemQuantity(t *testing.T) {
	log.Println("Item Quantity Check")

	itemQuantity := []string{"0", "100", "13", "we", "-100", "13.45", "-56.5"}

	for _, quantity := range itemQuantity {
		_, quantityError := validateItemPrice(quantity)
		if quantityError != nil {
			t.Errorf("Expected positive integer number only, got : %s", quantity)
		}
	}
}

func TestItemType(t *testing.T) {
	log.Println("Item Type Check")

	itemType := []string{"3", "1", "2", "-3", "4", "hi", "$#", "7.9", "0"}

	for _, types := range itemType {
		_, typesError := validateItemPrice(types)
		if typesError != nil {
			t.Errorf("Expected 1, 2 or 3 only, got : %s", types)
		}
	}
}
