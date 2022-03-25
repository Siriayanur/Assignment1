package model

import (
	"regexp"
	"strconv"

	"github.com/Siriayanur/Nuclei_Assignments/exceptions"
	"github.com/Siriayanur/Nuclei_Assignments/utils"
)

func validateItemName(itemName string) (string, error) {
	// Default itemName if blank
	if itemName == "" {
		return utils.DefaultItemName, nil
	}
	validName, _ := regexp.MatchString("^[a-zA-Z0-9_]*$", itemName)
	if validName {
		return itemName, nil
	}
	return itemName, exceptions.InvalidItemParameter("itemName", exceptions.ErrMap["itemName"])
}

func validateItemPrice(itemPrice string) (float64, error) {
	// Default itemPrice if blank
	if itemPrice == "" {
		return utils.DefaultItemPrice, nil
	}
	// Check positive float
	validFloat, _ := regexp.MatchString("^[^-]?([0-9]+([.][0-9]*)?|[.][0-9]+)$", itemPrice)
	if validFloat {
		validPrice, err := strconv.ParseFloat(itemPrice, 64)
		return utils.RoundFloat(validPrice), err
	}
	return utils.DefaultItemPrice, exceptions.InvalidItemParameter("itemPrice", exceptions.ErrMap["itemPrice"])
}

func validateItemQuantity(itemQuantity string) (int, error) {
	// Check blank
	if itemQuantity == "" {
		return utils.DefaultItemQuantity, nil
	}
	// Check positive integer
	validInt, _ := regexp.MatchString(`^\d*$`, itemQuantity)
	if validInt {
		validQuantity, err := strconv.Atoi(itemQuantity)
		return validQuantity, err
	}
	return utils.DefaultItemQuantity, exceptions.InvalidItemParameter("itemQuantity", exceptions.ErrMap["itemQuantity"])
}

func validateItemType(itemType string) (int, error) {
	if itemType != "" {
		// Check if it is in the specified range
		validType, _ := regexp.MatchString(`^[123]$`, itemType)
		if validType {
			validItemType, err := strconv.Atoi(itemType)
			return validItemType, err
		}
	}
	return 0, exceptions.InvalidItemParameter("itemType", exceptions.ErrMap["itemType"])
}
