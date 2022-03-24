package model

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/Siriayanur/Nuclei_Assignments/utils"
)

func validateItemName(itemName string) (string, error) {
	//Default itemName if blank
	if itemName == "" {
		return utils.ITEM_NAME, nil
	}
	validName, _ := regexp.MatchString("^[a-zA-Z0-9_]*$", itemName)
	if validName {
		return itemName, nil
	}
	return itemName, errors.New("Item Name should contain only Alphabets,Digits and Underscore\n")
}

func validateItemPrice(itemPrice string) (float64, error) {
	//Default itemPrice if blank
	if itemPrice == "" {
		return utils.ITEM_PRICE, nil
	}
	//Check positive float
	validFloat, _ := regexp.MatchString("^[^-]?([0-9]+([.][0-9]*)?|[.][0-9]+)$", itemPrice)
	if validFloat {
		validPrice, err := strconv.ParseFloat(itemPrice, 64)
		return validPrice, err
	} else {
		return utils.ITEM_PRICE, errors.New("Item Price should be a positive floating point number\n")
	}
}

func validateItemQuantity(itemQuantity string) (int, error) {
	// Check blank
	if itemQuantity == "" {
		return utils.ITEM_QUANTITY, nil
	}

	//Check positive integer
	validInt, _ := regexp.MatchString(`^0*[1-9]\d*$`, itemQuantity)
	if validInt {
		validQuantity, err := strconv.Atoi(itemQuantity)
		return validQuantity, err
	} else {
		return utils.ITEM_QUANTITY, errors.New("Item Quantity should be a positive integer\n")
	}
}

func validateItemType(itemType string) (int, error) {
	//check blank
	if itemType == "" {
		return 0, errors.New("Item Type cannot be blank")
	}

	//check
	validType, _ := regexp.MatchString(`^[123]$`, itemType)
	if validType {
		validItemType, err := strconv.Atoi(itemType)
		return validItemType, err
	} else {
		return 0, errors.New("Item Type should be one of these : 1/2/3\n")
	}
}
