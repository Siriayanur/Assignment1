package controller

import (
	"github.com/Siriayanur/Assignment1/model"
	"github.com/Siriayanur/Assignment1/utils"
)

func CalculateFinalPrice(item *model.Item) {
	mrp := 0.0
	item.SalesTax = utils.RoundFloat(getTax(item.Types, item.Price))
	mrp = float64(item.Quantity) * (item.Price + item.SalesTax)
	item.FinalPrice = utils.RoundFloat(mrp)
}

func getTax(itemType int, itemPrice float64) float64 {
	tax := 0.0

	switch itemType {
	case utils.RAW:
		tax = utils.BaseTax * itemPrice
	case utils.IMPORTED:
		tax = utils.BaseTax*itemPrice + (0.02 * (itemPrice + utils.BaseTax*itemPrice))
	case utils.MANUFACTURED:
		tax = utils.ImportDuty * itemPrice
		tax = tax + calculateSurcharge(tax, itemPrice)
	}
	return tax
}

func calculateSurcharge(amount float64, itemPrice float64) float64 {
	if amount <= 100 {
		return utils.SurchargeLevel1
	} else if amount <= 200 {
		return utils.SurchargeLevel2
	} else {
		return utils.SurchargeLevel3 * (itemPrice + (itemPrice * utils.ImportDuty))
	}
}
