package model

import (
	"strconv"

	"github.com/Siriayanur/Nuclei_Assignments/utils"
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
	return "Item Type : " + strconv.Itoa(item.Types)
}

func GetItem(itemName string, itemPrice string, itemQuantity string, itemType string) (Item, string) {

	var itemErrors = ""

	validItemName, itemNameError := validateItemName(itemName)
	validItemPrice, itemPriceError := validateItemPrice(itemPrice)
	validItemQuantity, itemQuantityError := validateItemQuantity(itemQuantity)
	validItemType, itemTypeError := validateItemType(itemType)

	if itemNameError != nil {
		itemErrors += itemNameError.Error()
	}
	if itemPriceError != nil {
		itemErrors += itemPriceError.Error()
	}
	if itemQuantityError != nil {
		itemErrors += itemQuantityError.Error()
	}
	if itemTypeError != nil {
		itemErrors += itemTypeError.Error()
	}

	item := Item{validItemName, validItemPrice, validItemQuantity, validItemType - 1, utils.SALES_TAX, utils.FINAL_PRICE}

	return item, itemErrors
}
