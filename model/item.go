package model

import (
	utils "github.com/Siriayanur/Nuclei_Assignments/utils"
)

type Item struct {
	Name       string
	Price      float64
	Quantity   int
	Types      int
	SalesTax   float64
	FinalPrice float64
}

func GetItem(itemName string, itemPrice float64, itemQuantity int, itemType int) Item {

	if itemName == "" {
		itemName = utils.ITEM_NAME
	}
	if itemPrice == 0 {
		itemPrice = utils.ITEM_PRICE
	}
	if itemQuantity == 0 {
		itemQuantity = utils.ITEM_QUANTITY
	}

	item := Item{itemName, itemPrice, itemQuantity, itemType, utils.SALES_TAX, utils.FINAL_PRICE}
	return item
}

func GetItemType(item Item) string {
	switch item.Types {
	case utils.RAW:
		return "Raw Item"
	case utils.IMPORTED:
		return "Imported Item"
	case utils.MANUFACTURED:
		return "Manufactured Item"
	}
	return ""
}
