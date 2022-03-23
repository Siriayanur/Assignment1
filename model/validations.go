package model

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/Siriayanur/Nuclei_Assignments/utils"
)

func validateItemName(itemName string) (error, string) {
	//Default itemName if blank
	if itemName == "" {
		return nil, utils.ITEM_NAME
	}
	validName, _ := regexp.MatchString("^[a-zA-Z0-9_]*$", itemName)
	if validName {
		return nil, itemName
	}
	return errors.New("Item Name should contain only Alphabets,Digits and Underscore\n"), itemName
}

func validateItemPrice(itemPrice string) (error, float64) {
	//Default itemPrice if blank
	if itemPrice == "" {
		return nil, utils.ITEM_PRICE
	}
	//Check positive float
	validFloat, _ := regexp.MatchString("^[^-]?([0-9]+([.][0-9]*)?|[.][0-9]+)$", itemPrice)
	if validFloat {
		validPrice, err := strconv.ParseFloat(itemPrice, 64)
		return err, validPrice
	} else {
		return errors.New("Item Price should be a positive floating point number\n"), utils.ITEM_PRICE
	}
}

func validateItemQuantity(itemQuantity string) (error, int) {
	// Check blank
	if itemQuantity == "" {
		return nil, utils.ITEM_QUANTITY
	}

	//Check positive integer
	validInt, _ := regexp.MatchString(`^0*[1-9]\d*$`, itemQuantity)
	if validInt {
		validQuantity, err := strconv.Atoi(itemQuantity)
		return err, validQuantity
	} else {
		return errors.New("Item Quantity should be a positive integer\n"), utils.ITEM_QUANTITY
	}
}

func validateItemType(itemType string) (error, int) {
	//check blank
	if itemType == "" {
		return errors.New("Item Type cannot be blank"), -1
	}

	//check
	validType, _ := regexp.MatchString(`^[123]$`, itemType)
	if validType {
		validItemType, err := strconv.Atoi(itemType)
		return err, validItemType
	} else {
		return errors.New("Item Type should be one of these : 1/2/3\n"), -1
	}
}
