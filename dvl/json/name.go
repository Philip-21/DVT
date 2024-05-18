package json

import (
	"encoding/json"

	"github.com/pkg/errors"
)


// validate name, returns an error if any and a JSON string format to be used for  specific purpose
func ValidateNameToString(name any) (string, error) {
	nameStr, ok := name.(string)
	if !ok {
		return "", errors.New("unable to convert to string")
	}
	jsonData, err := json.Marshal(nameStr)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil

}

// validates name and returns []byte, this give more flexibility,
// allowing you to work with the JSON data in its raw binary form, using []byte
// approach provides more control over how the JSON is handled
func ValidateNameToBytes(name any) ([]byte, error) {
	nameStr, ok := name.(string)
	if !ok {
		return nil, errors.New("unable to convert to string")
	}
	jsonData, err := json.Marshal(nameStr)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
