package model

import (
	"strconv"

	"github.com/Siriayanur/Assignment1/exceptions"
	"github.com/Siriayanur/Assignment1/utils"
)

type Item struct {
	Name       string
	Price      float64
	Quantity   int
	Types      int
	SalesTax   float64
	FinalPrice float64
}

func (item Item) GetItemType() string {
	return "Item Type : " + strconv.Itoa(item.Types+1)
}

func GetItem(itemName string, itemPrice string, itemQuantity string, itemType string) (Item, string) {
	exceptions.CreateErrorStatementsForItemParameters()
	validItemName, itemNameError := validateItemName(itemName)
	if itemNameError != nil {
		return Item{}, itemNameError.Error()
	}

	validItemPrice, itemPriceError := validateItemPrice(itemPrice)
	if itemPriceError != nil {
		return Item{}, itemPriceError.Error()
	}

	validItemQuantity, itemQuantityError := validateItemQuantity(itemQuantity)
	if itemQuantityError != nil {
		return Item{}, itemQuantityError.Error()
	}

	validItemType, itemTypeError := validateItemType(itemType)
	if itemTypeError != nil {
		return Item{}, itemTypeError.Error()
	}

	item := Item{validItemName, validItemPrice, validItemQuantity, validItemType - 1, utils.DefaultPrice, utils.DefaultPrice}

	return item, ""
}
