package controller

import (
	"fmt"
	"strconv"

	"github.com/Siriayanur/Nuclei_Assignments/model"
	"github.com/Siriayanur/Nuclei_Assignments/utils"
)

func CalculateFinalPrice(item *model.Item) {
	itemPrice, priceError := strconv.ParseFloat(item.Price, 64)
	itemQuantity, quantityError := strconv.Atoi(item.Quantity)
	mrp := 0.0
	if priceError != nil {
		fmt.Println(priceError.Error())
	}
	if quantityError != nil {
		fmt.Println(quantityError.Error())
	}
	item.SalesTax = getTax(item.Types, itemPrice)

	mrp = float64(itemQuantity) * (itemPrice + item.SalesTax)
	item.FinalPrice = mrp
}

func getTax(itemType int, itemPrice float64) float64 {
	tax := 0.0

	switch itemType {
	case utils.RAW:
		tax = utils.BASE_TAX * itemPrice
	case utils.IMPORTED:
		tax = utils.BASE_TAX*itemPrice + (0.02 * (itemPrice + utils.BASE_TAX*itemPrice))
	case utils.MANUFACTURED:
		tax = utils.IMPORT_DUTY * itemPrice
		tax = tax + CalculateSurcharge(tax)
	}
	return tax
}
