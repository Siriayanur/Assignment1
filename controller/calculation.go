package controller

import (
	"github.com/Siriayanur/Nuclei_Assignments/model"
	"github.com/Siriayanur/Nuclei_Assignments/utils"
)

func CalculateFinalPrice(item *model.Item) {
	mrp := 0.0

	item.SalesTax = getTax(item.Types, item.Price)
	mrp = float64(item.Quantity) * (item.Price + item.SalesTax)

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
		tax = tax + calculateSurcharge(tax, itemPrice)
	}
	return tax
}

func calculateSurcharge(amount float64, itemPrice float64) float64 {

	if amount <= 100 {
		return utils.SURCHARGE_LEVEL1
	} else if amount <= 200 {
		return utils.SURCHARGE_LEVEL2
	} else {
		return utils.SURCHARGE_LEVEL3 * (itemPrice + (itemPrice * utils.IMPORT_DUTY))
	}

}
