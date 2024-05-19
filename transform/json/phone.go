package json

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// validate phone number based on country code
// NGA, USA, ENG, GER, IND
// returns  a JSON string 
func PhoneToString(phone any) (string, error) {
	phonestr, ok := phone.(string)
	if !ok {
		return "", fmt.Errorf("unable to convert to string")
	}
	if len(phonestr) < 10 || phonestr[0] != '+' {
		return "", fmt.Errorf("invalid phonestr field format, ensure you have + added at the beginning of the phone number")
	}
	if len(phonestr) < 10 || phonestr[0] != '+' {
		return "", fmt.Errorf("invalid phonestr field format, ensure you have + added at the beginning of the phone number")
	}
	countryCodeNGA := `^\+234-\d{10}$`
	countryCodeUSA := `^\+1-\d{10}$`
	countryCodeGer := `^\+44-\d{10}$`
	countryCodeENG := `\+49-\d{10}$`
	countryCodeIND := `\+91-\d{10}$`
	var regEXP *regexp.Regexp

	switch {
	case regexp.MustCompile(countryCodeNGA).MatchString(phonestr):
		regEXP = regexp.MustCompile(countryCodeNGA)

	case regexp.MustCompile(countryCodeUSA).MatchString(phonestr):
		regEXP = regexp.MustCompile(countryCodeUSA)

	case regexp.MustCompile(countryCodeGer).MatchString(phonestr):
		regEXP = regexp.MustCompile(countryCodeGer)

	case regexp.MustCompile(countryCodeENG).MatchString(phonestr):
		regEXP = regexp.MustCompile(countryCodeENG)

	case regexp.MustCompile(countryCodeIND).MatchString(phonestr):
		regEXP = regexp.MustCompile(countryCodeIND)

	default:
		return "", fmt.Errorf("invalid country code: %s, valid country codes are +234, +1,+44, +49, +91", phonestr)
	}
	if !regEXP.MatchString(phonestr) {
		return "", fmt.Errorf("invalid country code: %s, valid country codes are +234, +1,+44, +49, +91", phonestr)
	}

	jsonData, err := json.Marshal(phonestr)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// validate phone number based on country code
// NGA, USA, ENG, GER, IND
// returns a []byte, this give more flexibility,
// allowing you to work with the JSON data in its raw binary form, using []byte
// approach provides more control over how the JSON is handled
func PhoneToBytes(phone any) ([]byte, error) {
	phonestr, ok := phone.(string)
	if !ok {
		return nil, fmt.Errorf("unable to convert to string")
	}
	if len(phonestr) < 10 || phonestr[0] != '+' {
		return nil, fmt.Errorf("invalid phonestr field format, ensure you have + added at the beginning of the phone number")
	}
	if len(phonestr) < 10 || phonestr[0] != '+' {
		return nil, fmt.Errorf("invalid phonestr field format, ensure you have + added at the beginning of the phone number")
	}
	countryCodeNGA := `^\+234-\d{10}$`
	countryCodeUSA := `^\+1-\d{10}$`
	countryCodeGer := `^\+44-\d{10}$`
	countryCodeENG := `\+49-\d{10}$`
	countryCodeIND := `\+91-\d{10}$`
	var regEXP *regexp.Regexp

	switch {
	case regexp.MustCompile(countryCodeNGA).MatchString(phonestr):
		regEXP = regexp.MustCompile(countryCodeNGA)

	case regexp.MustCompile(countryCodeUSA).MatchString(phonestr):
		regEXP = regexp.MustCompile(countryCodeUSA)

	case regexp.MustCompile(countryCodeGer).MatchString(phonestr):
		regEXP = regexp.MustCompile(countryCodeGer)

	case regexp.MustCompile(countryCodeENG).MatchString(phonestr):
		regEXP = regexp.MustCompile(countryCodeENG)

	case regexp.MustCompile(countryCodeIND).MatchString(phonestr):
		regEXP = regexp.MustCompile(countryCodeIND)

	default:
		return nil, fmt.Errorf("invalid country code: %s, valid country codes are +234, +1,+44, +49, +91", phonestr)
	}
	if !regEXP.MatchString(phonestr) {
		return nil, fmt.Errorf("invalid country code: %s, valid country codes are +234, +1,+44, +49, +91", phonestr)
	}

	jsonData, err := json.Marshal(phonestr)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
