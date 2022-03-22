package controller

import (
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
