package data

import (
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
)

func ValidateEmail(email any) (any, bool, error) {
	emailstr, ok := email.(string)
	if !ok {
		return nil, false, errors.New("unable to conveert to string")
	}
	if !govalidator.IsEmail(emailstr) {

		errMsg := fmt.Sprintf("invalid email format %s", emailstr)
		return nil, false, errors.New(errMsg)
	}
	return emailstr, true, nil
}
