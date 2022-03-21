package controller

import (
	"github.com/Siriayanur/Nuclei_Assignments/model"
	"github.com/Siriayanur/Nuclei_Assignments/utils"
)

func CalculateFinalPrice(item *model.Item) {

	mrp := 0.0
	tax := getTax(item.Types, item.Price)
	item.SalesTax = tax

	mrp = float64(item.Quantity) * (item.Price + tax)
	item.FinalPrice = mrp

}

func ValidateItemType(itemType int) bool {
	return (itemType >= 0) && (itemType <= 2)
}

func getTax(itemType int, itemPrice float64) float64 {
	tax := 0.0
	var cs TaxCalc

	switch itemType {
	case utils.RAW:
		tax = utils.BASE_TAX * itemPrice
	case utils.IMPORTED:
		tax = utils.BASE_TAX*itemPrice + (0.02 * (itemPrice + utils.BASE_TAX*itemPrice))
	case utils.MANUFACTURED:
		tax = utils.IMPORT_DUTY * itemPrice
		tax = tax + cs.CalculateSurcharge(tax)
	}
	return tax
}
