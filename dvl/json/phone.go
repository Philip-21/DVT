package json

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// validate phone number based on country code
// NGA, USA, ENG, GER, IND
// and returns an error if
func ValidatePhone(phone any) error {
	phonestr, ok := phone.(string)
	if !ok {
		return fmt.Errorf("unable to convert to string")
	}
	if len(phonestr) < 10 || phonestr[0] != '+' {
		return fmt.Errorf("invalid phonestr field format, ensure you have + added at the beginning of the phone number")
	}
	if len(phonestr) < 10 || phonestr[0] != '+' {
		return fmt.Errorf("invalid phonestr field format, ensure you have + added at the beginning of the phone number")
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
		return fmt.Errorf("invalid country code: %s, valid country codes are +234, +1,+44, +49, +91", phonestr)
	}
	if !regEXP.MatchString(phonestr) {
		return fmt.Errorf("invalid country code: %s, valid country codes are +234, +1,+44, +49, +91", phonestr)
	}

	_, err := json.Marshal(phonestr)
	if err != nil {
		return err
	}
	return nil
}

// validate phone number based on country code
// NGA, USA, ENG, GER, IND
// returns an error if any
// and a JSON string format to be used for  specific purpose
func ValidatePhoneToString(phone any) (string, error) {
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
// returns an error if any,
// and returns []byte, this give more flexibility,
// allowing you to work with the JSON data in its raw binary form, using []byte
// approach provides more control over how the JSON is handled
func ValidatePhoneToBytes(phone any) ([]byte, error) {
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

// func ValidatePhoneToBytes(phone any) ([]byte, error) {
// 	phonestr, ok := phone.(string)
// 	if !ok {
// 		return nil, fmt.Errorf("unable to convert to string")
// 	}
// 	if len(phonestr) < 10 || phonestr[0] != '+' {
// 		return nil, fmt.Errorf("invalid phone field format, ensure you have + added at the beginning of the phone number")
// 	}
// 	countryCodeNGA := phonestr[0:4]
// 	countryCodeUSA := phonestr[0:2]
// 	countryCodeENGGER := phonestr[0:3]
// 	var pattern string

// 	if countryCodeNGA == "+234" {
// 		// Handle Nigeria case
// 		pattern = `^\d{10}$`
// 		hyphenPos := -1
// 		for i := 0; i < len(phonestr); i++ {
// 			if phonestr[i] == '-' {
// 				hyphenPos = i
// 				break
// 			}
// 		}
// 		if hyphenPos == -1 {
// 			return nil, fmt.Errorf("invalid phone format, missing - after %s", countryCodeNGA)
// 		}
// 		phonestrNumberNigeria := phonestr[hyphenPos+1:]
// 		match, _ := regexp.MatchString(pattern, phonestrNumberNigeria)
// 		if !match {
// 			return nil, errors.New("invalid phone number format")
// 		}
// 	} else if countryCodeUSA == "+1" {
// 		pattern = `^\d{10}$` // USA
// 		hyphenPos := -1
// 		for i := 0; i < len(phonestr); i++ { //n.b i++ searches through the array to find -
// 			if phonestr[i] == '-' {
// 				hyphenPos = i
// 				break
// 			}
// 		}
// 		if hyphenPos == -1 {
// 			return nil, fmt.Errorf("invalid phone format, missing - after %s", countryCodeUSA)
// 		}
// 		phonestrNumberUSA := phonestr[hyphenPos+1:]
// 		match, _ := regexp.MatchString(pattern, phonestrNumberUSA)
// 		if !match {
// 			return nil, errors.New("invalid phone number format")
// 		}
// 	} else if countryCodeENGGER == "+44" {
// 		pattern = `^\d{10}$` // England
// 		hyphenPos := -1
// 		for i := 0; i < len(phonestr); i++ {
// 			if phonestr[i] == '-' { // '-' extract a single charac
// 				hyphenPos = i
// 				break
// 			}
// 		}
// 		if hyphenPos == -1 {
// 			return nil, fmt.Errorf("invalid phone format, missing - after %s", countryCodeENGGER)
// 		}
// 		phonestrNumberEngland := phonestr[hyphenPos+1:]
// 		match, _ := regexp.MatchString(pattern, phonestrNumberEngland)
// 		if !match {
// 			return nil, errors.New("invalid phone number format for")
// 		}

// 	} else if countryCodeENGGER == "+49" {
// 		pattern = `^\d{10}$` // Germany
// 		hyphenPos := -1
// 		for i := 0; i < len(phonestr); i++ {
// 			if phonestr[i] == '-' {
// 				hyphenPos = i
// 				break
// 			}
// 		}
// 		if hyphenPos == -1 {
// 			return nil, fmt.Errorf("invalid phone format, missing - after %s", countryCodeENGGER)
// 		}
// 		phonestrNumberGer := phonestr[hyphenPos+1:]
// 		match, _ := regexp.MatchString(pattern, phonestrNumberGer)
// 		if !match {
// 			return nil, fmt.Errorf("invalid phone number format")
// 		}
// 	} else {
// 		return nil, fmt.Errorf("invalid country code: %s, valid country codes are +234, +1,+44, +49, +91", phonestr)
// 	}
// 	jsonData, err := json.Marshal(phonestr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return jsonData, nil
// }
