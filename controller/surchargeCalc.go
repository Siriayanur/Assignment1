package controller

import "github.com/Siriayanur/Nuclei_Assignments/utils"

type TaxCalc interface {
	CalculateSurcharge(float64) float64
	getTax(int, float64) float64
}

func CalculateSurcharge(amount float64) float64 {

	if amount <= 100 {
		return utils.SURCHARGE_LEVEL1
	} else if amount <= 200 {
		return utils.SURCHARGE_LEVEL2
	} else {
		return utils.SURCHARGE_LEVEL3
	}

}
