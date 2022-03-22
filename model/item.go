package model

import (
	"regexp"

	"github.com/Siriayanur/Nuclei_Assignments/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Item struct {
	Name       string
	Price      string
	Quantity   string
	Types      int
	SalesTax   float64
	FinalPrice float64
	Error      string
}

func (item Item) GetItemType() string {
	switch item.Types {
	case utils.RAW:
		return "Raw Item"
	case utils.IMPORTED:
		return "Imported Item"
	case utils.MANUFACTURED:
		return "Manufactured Item"
	default:
		return ""
	}
}

func GetItem(itemName string, itemPrice string, itemQuantity string, itemType int) Item {

	if itemName == "" {
		itemName = utils.ITEM_NAME
	}
	if itemPrice == "" {
		itemPrice = utils.ITEM_PRICE
	}
	if itemQuantity == "" {
		itemQuantity = utils.ITEM_QUANTITY
	}

	item := Item{itemName, itemPrice, itemQuantity, itemType, utils.SALES_TAX, utils.FINAL_PRICE, ""}
	err := item.ValidateItem()

	//Check if validation error
	if err != nil {
		item.Error = err.Error()
	}
	return item
}

func (item Item) ValidateItem() error {

	return validation.ValidateStruct(&item,
		validation.Field(&item.Name, validation.Required, validation.Match(regexp.MustCompile("^[A-Za-z0-9_]*$"))),
		validation.Field(&item.Price, validation.Required, is.Float),
		validation.Field(&item.Quantity, validation.Required, is.Digit),
		validation.Field(&item.Types, validation.Required, validation.In(0, 1, 2)),
	)
}
