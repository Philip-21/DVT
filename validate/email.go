package validate

import (
	"encoding/json"
	
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/asaskevich/govalidator"
)

// validate an email, returns an error if any
func ValidateEmail(email any) error {
	emailStr, ok := email.(string)
	if !ok {
		return errors.New("unable to covert to string")
	}
	if !govalidator.IsEmail(emailStr) {
		errMsg := fmt.Sprintf("invalid email format %s", emailStr)
        return errors.New(errMsg)
	}
	_, err := json.Marshal(email)
	if err != nil {
		return err
	}
	return nil
}

// validate an email, returns an error if any and a JSON string format to be used for  specific purpose
func ValidateEmailToString(email any) (string, error) {
	emailStr, ok := email.(string)
	if !ok {
		return "", fmt.Errorf("invalid email format, expected a string")
	}
	if !govalidator.IsEmail(emailStr) {
		return "", fmt.Errorf("invalid email format %s", email)
	}
	jsonData, err := json.Marshal(email)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// validates email and returns []byte, this give more flexibility,
// allowing you to work with the JSON data in its raw binary form, using []byte
// approach provides more control over how the JSON is handled
func ValidateEmailtoBytes(email any) ([]byte, error) {
	emailStr, ok := email.(string)
	if !ok {
		return nil, fmt.Errorf("invalid email format, expected a string")
	}
	if !govalidator.IsEmail(emailStr) {
		log.Println("Invalid email format")
		return nil, fmt.Errorf("invalid email format %s", email)
	}
	jsonData, err := json.Marshal(email)
	if err != nil {
		return nil, fmt.Errorf("invalid email format %s", email)
	}
	return jsonData, nil
}
