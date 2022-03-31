package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestItemNameValidCases(t *testing.T) {
	itemNames := []string{"HeLlO", "PEN", "1234", "_HAT_", ""}
	for _, name := range itemNames {
		_, nameError := validateItemName(name)
		require.Nil(t, nameError)
	}
}
func TestItemNameInvalidCases(t *testing.T) {
	itemNames := []string{"hi@12", "$$@", "--", "he/<>1"}
	for _, name := range itemNames {
		_, nameError := validateItemName(name)
		require.NotNil(t, nameError)
	}
}
func TestItemPriceValidCases(t *testing.T) {
	itemPrices := []string{"120.00", "120", "34.333"}
	for _, price := range itemPrices {
		_, priceError := validateItemPrice(price)
		require.Nil(t, priceError)
	}
}
func TestItemPriceInvalidCases(t *testing.T) {
	itemPrices := []string{".", "uhmm", "120.abc", "-34.333", "-90", "@#"}
	for _, price := range itemPrices {
		_, priceError := validateItemPrice(price)
		require.NotNil(t, priceError)
	}
}
func TestItemQuantityValidCases(t *testing.T) {
	itemQuantity := []string{"0", "100", "13"}
	for _, quantity := range itemQuantity {
		_, quantityError := validateItemQuantity(quantity)
		require.Nil(t, quantityError)
	}
}
func TestItemQuantityInvalidCases(t *testing.T) {
	itemQuantity := []string{"we", "-100", "13.45", "-56.5"}
	for _, quantity := range itemQuantity {
		_, quantityError := validateItemQuantity(quantity)
		require.NotNil(t, quantityError)
	}
}
func TestItemTypeValidCases(t *testing.T) {
	itemType := []string{"3", "1", "2"}
	for _, types := range itemType {
		_, typesError := validateItemType(types)
		require.Nil(t, typesError)
	}
}
func TestItemTypeInvalidCases(t *testing.T) {
	itemType := []string{"-3", "4", "hi", "$#", "7.9", "0"}
	for _, types := range itemType {
		_, typesError := validateItemType(types)
		require.NotNil(t, typesError)
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
		require.Equal(t, "", err)
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
		require.NotEqual(t, "", err)
	}
}
