package transform

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type TransformDateTime struct {
	Regexps map[string]*regexp.Regexp
}

type TransformPhone struct {
	Countrycode map[string]*regexp.Regexp
}

// NewDateTimeTransform Initialize the  TransformDateTime struct
func NewDateTimeTransform() *TransformDateTime {
	return &TransformDateTime{
		Regexps: map[string]*regexp.Regexp{
			"gmt": regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+00:00$`), // "2023-07-10T16:22:41+00:00"
			"wat": regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+01:00$`), // "2023-07-10T16:22:41+01:00"
			"cet": regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+01:00$`), // "2023-07-10T16:22:41+01:00"
			"eet": regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+02:00$`), // "2023-07-10T16:22:41+02:00"
			"cst": regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\-06:00$`), // "2023-07-10T16:22:41-06:00"
		},
	}
}

// NewValidPhone initializes the TransformPhone struct
func NewValidPhone() *TransformPhone {
	countryCodeNGA := `^\+234-\d{10}$`
	countryCodeUSA := `^\+1-\d{10}$`
	countryCodeGER := `^\+44-\d{10}$`
	countryCodeENG := `\+49-\d{10}$`
	countryCodeIND := `\+91-\d{10}$`
	return &TransformPhone{
		Countrycode: map[string]*regexp.Regexp{
			"+234": regexp.MustCompile(countryCodeNGA),
			"+1":   regexp.MustCompile(countryCodeUSA),
			"+44":  regexp.MustCompile(countryCodeENG),
			"+49":  regexp.MustCompile(countryCodeGER),
			"+91":  regexp.MustCompile(countryCodeIND),
		},
	}
}

// validate phone number based on country code
// NGA, USA, ENG, GER, IND
// returns  a JSON string object
func (t *TransformPhone) PhoneToString(phone any) (string, error) {
	for _, re := range t.Countrycode {
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
		if re.MatchString(phonestr) {
			jsonBytes, err := json.Marshal(phonestr)
			if err != nil {
				return "", err
			}
			jsonString := string(jsonBytes)
			return jsonString, nil
		}
	}
	phonestr, ok := phone.(string)
	if !ok {
		return "", fmt.Errorf("unable to convert to string")
	}
	return "", fmt.Errorf("invalid country code: %s, valid country codes are +234, +1,+44, +49, +91", phonestr)
}

// validate phone number based on country code
// NGA, USA, ENG, GER, IND
// to a []byte, this give more flexibility,
// allowing you to work with the JSON data in its raw binary form, using []byte
// approach provides more control over how the JSON is handled
func (t *TransformPhone) PhoneToBytes(phone any) ([]byte, error) {
	for _, re := range t.Countrycode {
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
		if re.MatchString(phonestr) {
			jsonBytes, err := json.Marshal(phonestr)
			if err != nil {
				return nil, err
			}
			return jsonBytes, nil
		}
	}
	return nil, fmt.Errorf("invalid country code: %s, valid country codes are +234, +1,+44, +49, +91", phone)

}

// validates a datetime to a []byte type , this give more flexibility,
// allowing you to work with the JSON data in its raw binary form, using []byte
// approach provides more control over how the JSON is handled
func (t *TransformDateTime) DateTimeToBytes(timestamp string) ([]byte, error) {
	for _, re := range t.Regexps {
		if re.MatchString(timestamp) {
			jsonBytes, err := json.Marshal(timestamp)
			if err != nil {
				return nil, err
			}
			return jsonBytes, nil
		}
	}
	return nil, fmt.Errorf("invalid datetime format: %s", timestamp)
}

// validate's and converts datetime to a string JSON object.
func (t *TransformDateTime) DateTimeToString(timestamp string) (string, error) {
	for _, re := range t.Regexps {
		if re.MatchString(timestamp) {
			jsonBytes, err := json.Marshal(timestamp)
			if err != nil {
				return "", err
			}
			return string(jsonBytes), nil
		}
	}
	return "", fmt.Errorf("invalid datetime format: %s", timestamp)
}
