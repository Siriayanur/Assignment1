package controller

import (
	"github.com/Siriayanur/Nuclei_Assignments/model"
	"github.com/Siriayanur/Nuclei_Assignments/utils"
)

func CalculateSurcharge(amount float64) float64 {

	if amount <= 100 {
		return utils.SURCHARGE_LEVEL1
	} else if amount <= 200 {
		return utils.SURCHARGE_LEVEL2
	} else {
		return utils.SURCHARGE_LEVEL3
	}

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
