package exceptions

import (
	"errors"
	"fmt"
)

var ErrMap = make(map[string]string)

var errInvalidItemParameter = errors.New("invalid item parameter :: ")

func CreateErrorStatementsForItemParameters() {
	ErrMap["itemName"] = " contain only alphabets,digits and underscore "
	ErrMap["itemPrice"] = " be a positive floating point number  "
	ErrMap["itemQuantity"] = " be a positive integer "
	ErrMap["itemType"] = " not be blank ,be one of these : 1/2/3 "
}

func InvalidItemParameter(parameter string, statement string) error {
	return fmt.Errorf("%w :: %s must %s", errInvalidItemParameter, parameter, statement)
}
