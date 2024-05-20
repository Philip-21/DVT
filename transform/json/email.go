package json

import (
	"encoding/json"

	"fmt"
	"log"

	"github.com/asaskevich/govalidator"
)

// validate's and converts  an email to  a string JSON object.
func EmailToString(email any) (string, bool, error) {
	emailStr, ok := email.(string)
	if !ok {
		return "", false, fmt.Errorf("invalid email format, expected a string")
	}
	if !govalidator.IsEmail(emailStr) {
		return "", false, fmt.Errorf("invalid email format %s", email)
	}

	jsonData, err := json.Marshal(email)
	if err != nil {
		return "", false, err
	}
	return string(jsonData), true, nil
}

// validates an email to a []byte type , this give more flexibility,
// allowing you to work with the JSON data in its raw binary form, using []byte
// approach provides more control over how the JSON is handled
func EmailtoBytes(email any) ([]byte, error) {
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
