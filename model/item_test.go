package model

import (
	"fmt"
	"log"
	"testing"
)

func TestItemNameHappyFlow(t *testing.T) {

	fmt.Println("HAPPY FLOW - Item Name")

	itemNames := []string{"HeLlO", "PEN", "pencil", "_HAT_", "App12", "1200", ""}

	for _, name := range itemNames {
		validName, nameError := validateItemName(name)
		log.Println(validName, nameError)
	}

}

func TestItemNameBadFlow(t *testing.T) {

	fmt.Println("BAD FLOW - Item Name")

	itemName := "hi@12"
	validName, nameError := validateItemName(itemName)

	log.Println(validName, nameError)

	itemName = "$$@"
	validName, nameError = validateItemName(itemName)
	log.Println(validName, nameError)

}

func TestItemPriceHappyFlow(t *testing.T) {

	fmt.Println("HAPPY FLOW - Item Price")

	itemPrices := []string{"120.00", "120", "34.333", "."}

	for _, price := range itemPrices {
		validPrice, priceError := validateItemPrice(price)
		log.Println(validPrice, priceError)
	}

}

func TestItemPriceBadFlow(t *testing.T) {
	fmt.Println("BAD FLOW - Item Price")

	itemPrices := []string{"uhmm", "120.abc", "-34.333", "-90", "@#"}

	for _, price := range itemPrices {
		validPrice, priceError := validateItemPrice(price)
		log.Println(validPrice, priceError)
	}
}

func TestItemQuantityHappyFlow(t *testing.T) {
	fmt.Println("HAPPY FLOW - Item Quantity")

	itemQuantity := []string{"0", "100", "13"}

	for _, price := range itemQuantity {
		validQuantity, quantityError := validateItemQuantity(price)
		log.Println(validQuantity, quantityError)
	}
}

func TestItemQuantityBadFlow(t *testing.T) {
	fmt.Println("BAD FLOW - Item Quantity")

	itemQuantity := []string{"we", "-100", "13.45"}

	for _, price := range itemQuantity {
		validQuantity, quantityError := validateItemQuantity(price)
		log.Println(validQuantity, quantityError)
	}
}

func TestItemTypeHappyFlow(t *testing.T) {
	fmt.Println("HAPPY FLOW - Item Type")

	itemType := []string{"3", "1", "2"}

	for _, types := range itemType {
		validQuantity, quantityError := validateItemType(types)
		log.Println(validQuantity, quantityError)
	}
}
func TestItemTypeBadFlow(t *testing.T) {
	fmt.Println("BAD FLOW - Item Type")

	itemType := []string{"-3", "4", "hi", "$#"}

	for _, types := range itemType {
		validQuantity, quantityError := validateItemType(types)
		log.Println(validQuantity, quantityError)
	}
}
